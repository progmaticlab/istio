// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conformance

import (
	"fmt"
	"os"
	"path"
	"testing"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/environment"
	"istio.io/istio/pkg/test/framework/components/galley"
	"istio.io/istio/pkg/test/framework/components/istio"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/pkg/test/framework/label"
	compliance2 "istio.io/istio/tests/integration/conformance/pkg/conformance"
	constraint2 "istio.io/istio/tests/integration/conformance/pkg/conformance/constraint"
)

func TestCompliance(t *testing.T) {
	framework.Run(t, func(ctx framework.TestContext) {
		cases, err := loadCases()
		if err != nil {
			ctx.Fatalf("error loading test cases: %v", err)
		}

		gal := galley.NewOrFail(ctx, ctx, galley.Config{})

		for _, ca := range cases {
			tst := ctx.NewSubTest(ca.Metadata.Name)

			for _, lname := range ca.Metadata.Labels {
				l, ok := label.Find(lname)
				if !ok {
					ctx.Fatalf("label not found: %v", lname)
				}
				tst = tst.Label(l)
			}

			if ca.Metadata.Parallel {
				tst.RunParallel(getRunTestFn(gal, ca))
			} else {
				tst.Run(getRunTestFn(gal, ca))
			}
		}
	})
}

func getRunTestFn(gal galley.Instance, ca *compliance2.Test) func(framework.TestContext) {
	return func(ctx framework.TestContext) {
		match := true
	mainloop:
		for _, ename := range ca.Metadata.Environments {
			match = false
			for _, n := range environment.Names() {
				if n.String() == ename && n == ctx.Environment().EnvironmentName() {
					match = true
					break mainloop
				}
			}
		}

		if !match {
			ctx.Skipf("None of the expected environment(s) not found: %v", ca.Metadata.Environments)
		}

		if ca.Metadata.Skip {
			ctx.Skipf("Test is marked as skip")
		}

		if err := gal.ClearConfig(); err != nil {
			ctx.Fatalf("Error clearing config: %v", err)
		}

		ns := namespace.NewOrFail(ctx, ctx, "conv", true)

		if len(ca.Stages) == 1 {
			runStage(ctx, gal, ns, ca.Stages[0])
		} else {
			for i, s := range ca.Stages {
				ctx.NewSubTest(fmt.Sprintf("%d", i)).Run(func(ctx framework.TestContext) {
					runStage(ctx, gal, ns, s)
				})
			}
		}
	}
}

func runStage(ctx framework.TestContext, gal galley.Instance, ns namespace.Instance, s *compliance2.Stage) {
	i := s.Input
	gal.ApplyConfigOrFail(ctx, ns, i)

	if s.MCP != nil {
		validateMCPState(ctx, gal, ns, s)
	}

	// More and different types of validations can go here
}

func validateMCPState(ctx framework.TestContext, gal galley.Instance, ns namespace.Instance, s *compliance2.Stage) {
	p := constraint2.Params{
		Namespace: ns.Name(),
	}
	for _, coll := range s.MCP.Constraints {
		gal.WaitForSnapshotOrFail(ctx, coll.Name, func(actuals []*galley.SnapshotObject) error {
			for _, rangeCheck := range coll.Check {
				a := make([]interface{}, len(actuals))
				for i, item := range actuals {
					a[i] = item
					// Clear out for stable comparison.
					item.Metadata.CreateTime = nil
					item.Metadata.Annotations = nil
					item.Metadata.Version = ""
				}

				if err := rangeCheck.ValidateItems(a, p); err != nil {
					return err
				}
			}
			return nil
		})
	}

}

func loadCases() ([]*compliance2.Test, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	p := path.Join(wd, "cases")
	return compliance2.Load(p)
}

func TestMain(m *testing.M) {
	framework.
		NewSuite("compliance_test", m).
		SetupOnEnv(environment.Kube, istio.Setup(nil, nil)).
		Run()
}

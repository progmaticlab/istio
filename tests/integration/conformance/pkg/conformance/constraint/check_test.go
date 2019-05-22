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

package constraint

// TODO
//
// func TestCheck(t *testing.T) {
// 	cases := []struct {
// 		Input string
// 		Items []interface{}
// 		Err   bool
// 	}{
// 		{
// 			Input: `
// constraints:
//   - collection: col1
//     check:
//     - exactlyOne:
//       - select: bar
//         exists: true
//         then:
//           - select
// `,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		t.Run("", func(t *testing.T) {
// 			g := NewGomegaWithT(t)
// 			constrs, err := Parse([]byte(c.Input))
// 			g.Expect(err).To(BeNil())
//
// 			coll := constrs.Constraints[0].Check[0]
// 			err = coll.ValidateItems(c.Items)
// 			if c.Err {
// 				g.Expect(err).NotTo(BeNil())
// 			} else {
// 				g.Expect(err).To(BeNil())
// 			}
// 		})
// 	}
// }

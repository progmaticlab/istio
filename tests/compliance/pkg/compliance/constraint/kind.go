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

// Kind of constraint
type Kind string

//
// const (
// 	Collection Kind = "collection"
//
// 	// Any entry should match the constraint
// 	Any Kind = "Any"
//
// 	// ExactlyOne entry should match the constraint
// 	ExactlyOne Kind = "ExactlyOne"
//
// 	// Select an entry based on structpath
// 	Select Kind = "select"
//
// 	// Exists checks for existence, based on structpath
// 	Exact Kind = "exact"
//
// 	// Equals checks for equality based on structpath
// 	Equals Kind = "equals"
//
// 	// NotEquals checks for inequality based on structpath
// 	NotEquals Kind = "notEquals"
// )

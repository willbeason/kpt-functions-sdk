// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package language

import (
	"github.com/willbeason/typegen/pkg/definition"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func expectEqual(t *testing.T, expected, actual string) {
	t.Helper()
	if diff := cmp.Diff(strings.Split(expected, "\n"), strings.Split(actual, "\n")); diff != "" {
		t.Fatal(diff)
	}
}

func TestAlias_Print(t *testing.T) {
	testCases := []struct {
		name     string
		alias    definition.Alias
		expected string
	}{
		{
			name: "Empty object",
			alias: definition.Alias{
				Meta: definition.Meta{
					Name:        "Empty",
					Description: "An empty definition.",
				},
				Type: definition.Empty{},
			},
			expected: `// An empty definition.
export type Empty = object;`,
		},
		{
			name: "Number alias",
			alias: definition.Alias{
				Meta: definition.Meta{
					Name: "Quantity",
				},
				Type: definition.KnownPrimitives.Integer,
			},
			expected: `export type Quantity = number;`,
		},
		{
			name: "Map of array alias",
			alias: definition.Alias{
				Meta: definition.Meta{
					Name: "QuantitiesMap",
				},
				Type: definition.Map{Values: definition.Array{Items: definition.KnownPrimitives.Integer}},
			},
			expected: `export type QuantitiesMap = {[key: string]: number[]};`,
		},
	}

	for _, tc := range testCases {
		ts := TypeScript{
			RefObjects: make(map[definition.Ref]definition.Object),
		}
		t.Run(tc.name, func(t *testing.T) {
			// For Alias types, we expect identical code for both the class and the type.
			expectEqual(t, tc.expected, ts.PrintDefinition(tc.alias))
		})
	}
}

func TestModel_Print(t *testing.T) {
	testCases := []struct {
		name              string
		refObjects        map[definition.Ref]definition.Object
		model             definition.Object
		expected          string
		expectedInterface string
	}{
		{
			name: "Empty properties",
			model: definition.Object{
				Meta: definition.Meta{
					Name:        "Empty",
					Description: "An empty model.",
				},
			},
			expected: `// An empty model.
export class Empty {

}`,
		},
		{
			name: "One property",
			model: definition.Object{
				Meta: definition.Meta{
					Name: "Quantity",
				},
				Properties: map[string]definition.Property{
					"value": {
						Description: "The amount represented by the Quantity.",
						Type:        definition.KnownPrimitives.Integer,
						Required:    true,
					},
				},
			},
			expected: `export class Quantity {
  // The amount represented by the Quantity.
  public value: number;

  constructor(desc: Quantity) {
    this.value = desc.value;
  }
}`,
		},
		{
			name: "Two properties",
			model: definition.Object{
				Meta: definition.Meta{
					Name: "Quantity_v2",
				},
				Properties: map[string]definition.Property{
					"value": {
						Description: "The amount represented by the Quantity_v2.",
						Type:        definition.KnownPrimitives.Integer,
						Required:    true,
					},
					"unit": {
						Description: "The unit value is expressed in.",
						Type:        definition.KnownPrimitives.String,
					},
				},
			},
			expected: `export class Quantity_v2 {
  // The unit value is expressed in.
  public unit?: string;

  // The amount represented by the Quantity_v2.
  public value: number;

  constructor(desc: Quantity_v2) {
    this.unit = desc.unit;
    this.value = desc.value;
  }
}`,
		},
		{
			name: "Nested types",
			refObjects: map[definition.Ref]definition.Object{
				definition.Ref{
					Package: "api.io.v1",
					Name:    "Pod.Spec",
				}: {Properties: map[string]definition.Property{"required": {Required: true}}},
			},
			model: definition.Object{
				Meta: definition.Meta{
					Name:        "Pod",
					Package:     "api.io.v1",
					Description: "a complex Pod model",
				},
				Properties: map[string]definition.Property{
					"spec": {
						Type: definition.Ref{
							Package: "api.io.v1",
							Name:    "Pod.Spec",
						},
						Description: "The Pod specification.",
					},
				},
				NestedTypes: []definition.Object{
					{
						Meta: definition.Meta{
							Name:        "Spec",
							Package:     "api.io.v1",
							Namespace:   []string{"Pod"},
							Description: "The Pod specification.",
						},
						Properties: map[string]definition.Property{
							"restartStrategy": {
								Type: definition.Ref{
									Package: "api.io.v1",
									Name:    "Pod.Spec.RestartStrategy",
								},
								Description: "When to restart the Pod.",
								Required:    true,
							},
						},
						NestedTypes: []definition.Object{
							{
								Meta: definition.Meta{
									Name:        "RestartStrategy",
									Package:     "api.io.v1",
									Namespace:   []string{"Pod", "Spec"},
									Description: "When to restart the Pod.",
								},
							},
						},
					},
				},
			},
			expected: `// a complex Pod model
export class Pod {
  // The Pod specification.
  public spec?: Pod.Spec;
}

export namespace Pod {
  // The Pod specification.
  export class Spec {
    // When to restart the Pod.
    public restartStrategy: Pod.Spec.RestartStrategy;

    constructor(desc: Pod.Spec) {
      this.restartStrategy = desc.restartStrategy;
    }
  }

  export namespace Spec {
    // When to restart the Pod.
    export class RestartStrategy {

    }
  }
}`,
		},
	}

	for _, tc := range testCases {
		ts := TypeScript{
			RefObjects: tc.refObjects,
		}
		t.Run(tc.name, func(t *testing.T) {
			expectEqual(t, tc.expected, ts.PrintDefinition(tc.model))
		})
	}
}

func TestPropertiesTSTypes(t *testing.T) {
	testCases := []struct {
		name     string
		pkg      string
		property definition.NamedProperty
		expected string
	}{
		{
			name: "Print the field definition for types.go.",
			pkg:  "api.io.v1",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type: definition.Ref{
						Package: "api.io.v1",
						Name:    "Pod.PodSpec",
					},
					Required: true,
				},
			},
			expected: `// A description.
public spec: Pod.PodSpec;`,
		},
		{
			name: "Print the field definition in a different package.",
			pkg:  "api.io.v1",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type: definition.Ref{
						Package: "api.io.v2",
						Name:    "Pod.PodSpec",
					},
					Required: true,
				},
			},
			expected: `// A description.
public spec: apiIoV2.Pod.PodSpec;`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := PrintTSTypesField(tc.pkg, tc.property)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestPropertiesTSConstructor(t *testing.T) {
	testCases := []struct {
		name     string
		property definition.NamedProperty
		expected string
	}{
		{
			name: "primitive field",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type:        definition.KnownPrimitives.Integer,
					Required:    true,
				}},
			expected: `
this.spec = desc.spec;`,
		},
		{
			name: "optional primitive field",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type:        definition.KnownPrimitives.Integer,
				},
			},
			expected: `
this.spec = desc.spec;`,
		},
		{
			name: "ref field",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type:        definition.Ref{Name: "io.PodSpec"},
					Required:    true,
				},
			},
			expected: `
this.spec = desc.spec;`,
		},
		{
			name: "optional ref field",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type:        definition.Ref{Name: "io.PodSpec"},
				},
			},
			expected: `
this.spec = desc.spec;`,
		},
		{
			name: "array field",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type:        definition.Array{Items: definition.Ref{Name: "io.PodSpec"}},
					Required:    true,
				},
			},
			expected: `
this.spec = desc.spec;`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := TypeScript{
				RefObjects: make(map[definition.Ref]definition.Object),
			}
			actual := ts.PrintTSConstructorField("", tc.property)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestPropertiesTSInterfaces(t *testing.T) {
	testCases := []struct {
		name     string
		pkg      string
		property definition.NamedProperty
		expected string
	}{
		{
			name: "Print the field definition for interfaces.go.",
			pkg:  "api.io.v1",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type: definition.Ref{
						Package: "api.io.v1",
						Name:    "PodSpec",
					},
					Required: true,
				},
			},
			expected: `// A description.
spec: PodSpec;`,
		},
		{
			name: "Print the field definition from type in different package.",
			property: definition.NamedProperty{
				Name: "spec",
				Property: definition.Property{
					Description: "A description.",
					Type: definition.Ref{
						Package: "api.io.v2",
						Name:    "PodSpec",
					},
					Required: true,
				},
			},
			expected: `// A description.
spec: apiIoV2.PodSpec;`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := PrintTSInterfacesField(tc.pkg, tc.property)

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

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

package swagger

import (
	"testing"

	"github.com/willbeason/typegen/pkg/definition"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func TestParseDefinition(t *testing.T) {
	testCases := []struct {
		name      string
		modelName string
		input     map[string]interface{}
		expected  definition.Definition
	}{
		{
			name:      "Empty declaration",
			modelName: "io.v1.Empty",
			input:     map[string]interface{}{},
			expected: definition.Alias{
				Meta: definition.Meta{
					Name:    "Empty",
					Package: "io.v1",
				},
				Type: definition.Empty{},
			},
		},
		{
			name:      "Type reference",
			modelName: "io.v1alpha1.Pod",
			input: map[string]interface{}{
				"$ref": "#/definitions/io.v1.Pod",
			},
			expected: definition.Alias{
				Meta: definition.Meta{
					Name:    "Pod",
					Package: "io.v1alpha1",
				},
				Type: definition.Ref{
					Package: "io.v1",
					Name:    "Pod",
				},
			},
		},
		{
			name:      "Type alias",
			modelName: "io.v1.Quantity",
			input: map[string]interface{}{
				"type":   "integer",
				"format": "int32",
			},
			expected: definition.Alias{
				Meta: definition.Meta{
					Name:    "Quantity",
					Package: "io.v1",
				},
				Type: definition.KnownPrimitives.Integer,
			},
		},
		{
			name:      "Model with two fields",
			modelName: "io.v1.Pod",
			input: map[string]interface{}{
				"description": "a simple Pod model",
				"properties": map[string]interface{}{
					"podType": map[string]interface{}{
						"type": "string",
					},
					"spec": map[string]interface{}{
						"description": "unstructured spec field",
					},
				},
			},
			expected: definition.Object{
				Meta: definition.Meta{
					Name:        "Pod",
					Package:     "io.v1",
					Description: "a simple Pod model",
				},
				Properties: map[string]definition.Property{
					"podType": {
						Type: definition.KnownPrimitives.String,
					},
					"spec": {
						Type:        definition.Empty{},
						Description: "unstructured spec field",
					},
				},
			},
		},
		{
			name:      "Model with nested type with nested type",
			modelName: "io.v1.Pod",
			input: map[string]interface{}{
				"description": "a complex Pod model",
				"properties": map[string]interface{}{
					"spec": map[string]interface{}{
						"properties": map[string]interface{}{
							"restartStrategy": map[string]interface{}{
								"properties": map[string]interface{}{},
							},
						},
					},
				},
			},
			expected: definition.Object{
				Meta: definition.Meta{
					Name:        "Pod",
					Package:     "io.v1",
					Description: "a complex Pod model",
				},
				Properties: map[string]definition.Property{
					"spec": {
						Type: definition.Ref{
							Package: "io.v1",
							Name:    "Pod.Spec",
						},
					},
				},
				NestedTypes: []definition.Object{
					{
						Meta: definition.Meta{
							Name:      "Spec",
							Package:   "io.v1",
							Namespace: []string{"Pod"},
						},
						Properties: map[string]definition.Property{
							"restartStrategy": {
								Type: definition.Ref{
									Package: "io.v1",
									Name:    "Pod.Spec.RestartStrategy",
								},
							},
						},
						NestedTypes: []definition.Object{
							{
								Meta: definition.Meta{
									Name:      "RestartStrategy",
									Package:   "io.v1",
									Namespace: []string{"Pod", "Spec"},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := newParser()
			result := p.parseDefinition(tc.modelName, tc.input)

			if diff := cmp.Diff(tc.expected, result, cmpopts.EquateEmpty()); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

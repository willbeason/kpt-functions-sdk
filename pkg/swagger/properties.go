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
	"github.com/willbeason/typegen/pkg/definition"
	"github.com/willbeason/typegen/pkg/maps"
	"strings"
)



// parseProperties parses the []Properties defined by a Model.
//
// Returns the contained properties and nested type definitions.
func (p parser) parseProperties(definitionMeta definition.Meta, model map[string]interface{}) (map[string]definition.Property, []definition.Object) {
	requiredFields, _ := maps.GetStringArray("required", model)
	required := make(map[string]bool)
	for _, field := range requiredFields {
		required[field] = true
	}

	propertiesMap, hasProperties := maps.GetMap("properties", model)
	if !hasProperties {
		return nil, nil
	}

	properties := make(map[string]definition.Property)
	var nestedTypes []definition.Object
	for name := range propertiesMap {
		if isUnsupportedProperty(name) {
			continue
		}

		propertyMap := maps.GetRequiredMap(name, propertiesMap)

		description, _ := maps.GetString("description", propertyMap)

		var typ definition.Type
		if isObject(propertyMap) {
			// This property has "properties", so it is an object with a complex definition.
			propertyType := strings.Title(name)
			propertyDefinitionMeta := definition.Meta{
				Name:        propertyType,
				Package:     definitionMeta.Package,
				Namespace:   append(definitionMeta.Namespace, definitionMeta.Name),
				Description: description,
			}
			object := p.parseObject(propertyDefinitionMeta, propertyMap)
			nestedTypes = append(nestedTypes, object)
			typ = definition.Ref{
				Name:    strings.Join(append(propertyDefinitionMeta.Namespace, propertyType), "."),
				Package: definitionMeta.Package,
			}
		} else {
			typ = p.parseType(definitionMeta, propertyMap)
		}

		properties[name] = definition.Property{
			Type:        typ,
			Description: description,
			Required:    required[name],
		}
		nestedTypes = append(nestedTypes, typ.NestedTypes()...)
	}

	return properties, nestedTypes
}

// Excludes properties that are k8s extensions (e.g. x-kubernetes-*) since they cause issues when generating fields in e.g. TS.
// This currently only affects CRD definition like
// io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.JSONSchemaProps
// We can consider better handling this if it's a real use case.
func isUnsupportedProperty(name string) bool {
	return strings.Contains(name, "x-kubernetes-")
}

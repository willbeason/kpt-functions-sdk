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
	"fmt"
	"github.com/willbeason/typegen/pkg/definition"
	"github.com/willbeason/typegen/pkg/maps"
)


func isObject(m map[string]interface{}) bool {
	_, result := m["properties"]
	return result
}

// parseObject parses a model given its key in the definitions map and the map holding all of its information.
func (p parser) parseObject(meta definition.Meta, model map[string]interface{}) definition.Object {
	properties, nestedTypes := p.parseProperties(meta, model)

	o := definition.Object{
		Meta:        meta,
		NestedTypes: nestedTypes,
		Properties:  properties,
	}

	if gvks, containsGVK := model["x-kubernetes-group-version-kind"]; containsGVK {
		o.GroupVersionKinds = getGVKs(gvks)

		// The object declares GroupVersionKind, so require that apiVersion and kind exist even if they aren't declared.
		o.Properties["apiVersion"] = o.Properties["apiVersion"].WithRequired().WithType(definition.KnownPrimitives.String)
		o.Properties["kind"] = o.Properties["kind"].WithRequired().WithType(definition.KnownPrimitives.String)

		// There are types with GVK but no "metadata". These aren't usually recognized by the API Server, and are
		// usually required for ones that are.
		if metadata, hasMetadata := o.Properties["metadata"]; hasMetadata {
			if ref, isRef := metadata.Type.(definition.Ref); isRef {
				// TODO(b/142003702): Allow other ObjectMeta types.
				if ref.Name == "ObjectMeta" && ref.Package == "io.k8s.apimachinery.pkg.apis.meta.v1" {
					// Types don't usually declare "metadata" as required even though they actually are.
					properties["metadata"] = properties["metadata"].WithRequired()
					o.IsKubernetesObject = true
				}
			}
		}
	}

	if o.Package == "io.k8s.apimachinery.pkg.apis.meta.v1" && o.Name == "ObjectMeta" {
		// Normally either "name" or "generateName" must be defined in ObjectMeta. We don't support "generateName", so
		// we require "name".
		o.Properties["name"] = o.Properties["name"].WithRequired()
	}

	p.RefObjects[o.ToRef()] = o
	return o
}

func getGVKs(gvks interface{}) []definition.GroupVersionKind {
	var result []definition.GroupVersionKind
	gvksArray, ok := gvks.([]interface{})
	if !ok {
		panic(fmt.Sprintf("x-kubernetes-group-version-kind must be an array: %+v", gvks))
	}

	for _, gvkInterface := range gvksArray {
		gvkMap, isMap := gvkInterface.(map[string]interface{})
		if !isMap {
			panic(fmt.Sprintf("x-kubernetes-group-version-kind must be an array of maps: %+v", gvks))
		}
		result = append(result, definition.GroupVersionKind{
			Group:   maps.GetRequiredString("group", gvkMap),
			Version: maps.GetRequiredString("version", gvkMap),
			Kind:    maps.GetRequiredString("kind", gvkMap),
		})
	}
	return result
}

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


func (p parser) parseArray(definitionMeta definition.Meta, o map[string]interface{}) definition.Array {
	itemsMap := maps.GetRequiredMap("items", o)
	if isObject(itemsMap) {
		description, _ := maps.GetString("description", itemsMap)
		meta := definition.Meta{
			Name:        "Item",
			Package:     definitionMeta.Package,
			Namespace:   append(definitionMeta.Namespace, definitionMeta.Name),
			Description: description,
		}
		object := p.parseObject(meta, itemsMap)
		return definition.Array{
			Items: definition.Ref{
				Package: definitionMeta.Package,
				Name:    strings.Join(append(meta.Namespace, "Item"), "."),
			},
			Nested: []definition.Object{object},
		}
	}
	return definition.Array{
		Items: p.parseType(definitionMeta, itemsMap),
	}
}

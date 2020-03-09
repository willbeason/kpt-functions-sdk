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
	"strings"

	"github.com/willbeason/typegen/pkg/definition"
)

func refMatches(ref definition.Ref, filter []string) bool {
	fullName := ref.Name
	if ref.Package != "" {
		fullName = ref.Package + "." + fullName
	}
	if !strings.HasPrefix(fullName, filter[0]) {
		return false
	}

	idx := len(filter[0])
	for i := range filter {
		if i == 0 {
			continue
		}
		inc := strings.Index(fullName[idx:], filter[i])
		if inc == -1 {
			return false
		}
		idx += inc + len(filter[i])
	}
	return true
}

// FilterDefinitions returns the filtered subset of Definitions and their transitive dependencies.
func FilterDefinitions(filters []string, allPackages map[string][]definition.Definition) map[string][]definition.Definition {
	// Record all definitions by their reference.
	allDefinitions := make(map[definition.Ref]definition.Definition)
	// Record all package-level dependencies.
	allDependencies := make(map[definition.Ref]map[definition.Ref]bool)
	for _, definitions := range allPackages {
		for _, d := range definitions {
			ref := d.Metadata().ToRef()
			// Map reference to its definition.
			allDefinitions[ref] = d
			// Determine dependencies for this definition.
			allDependencies[ref] = make(map[definition.Ref]bool)
			imports := d.Imports()
			for _, i := range imports {
				allDependencies[ref][i] = true
			}
		}
	}

	// Record Refs matching filters.
	var refs []definition.Ref
	for _, definitions := range allPackages {
		for _, d := range definitions {
			for _, filter := range filters {
				ref := d.Metadata().ToRef()
				if refMatches(ref, strings.Split(filter, "*")) {
					refs = append(refs, ref)
					break
				}
			}
		}
	}

	// Determine set of included Definitions and their transitive dependencies.
	definitions := make(map[definition.Ref]definition.Definition)
	for i := 0; i < len(refs); i++ {
		ref := refs[i]
		if _, found := definitions[ref]; found {
			// We have already included this Definition.
			continue
		}
		if strings.Contains(ref.Name, ".") {
			// This is a ref to a nested class and already implicitly included.
			continue
		}

		// Include this definition
		definitions[ref] = allDefinitions[ref]

		// Mark all dependencies of this definition to be included.
		for dependency := range allDependencies[ref] {
			refs = append(refs, dependency)
		}
	}

	// Fill in included Definitions.
	result := make(map[string][]definition.Definition)
	for ref, d := range definitions {
		result[ref.Package] = append(result[ref.Package], d)
	}
	return result
}

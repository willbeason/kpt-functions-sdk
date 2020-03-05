package crds

import (
	"fmt"
	"github.com/willbeason/typegen/pkg/definition"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Parse parses a CRD into a definition.Definition.
func Parse(spec v1.CustomResourceDefinitionSpec) definition.Definition {
	// TODO(willbeason): Don't assume the first version is the only desired one.
	version := spec.Versions[0]
	schema := version.Schema.OpenAPIV3Schema
	def := definition.Object{
		Meta: definition.Meta{
			Name:        spec.Names.Kind,
			Package:     fmt.Sprintf("%s.%s", spec.Group, version.Name),
			Description: schema.Description,
		},
		Properties: make(map[string]definition.Property),
		IsKubernetesObject: true,
		GroupVersionKinds: gvks(spec.Group, spec.Names.Kind, spec.Versions),
	}

	// Assume the CRD contains properties.
	for name, property := range schema.Properties {
		def.Properties[name] = parseProperty(property)
	}
	for _, name := range schema.Required {
		p := def.Properties[name]
		p.Required = true
		def.Properties[name] = p
	}

	return def
}

func gvks(group string, kind string, versions []v1.CustomResourceDefinitionVersion) []definition.GroupVersionKind {
	var result []definition.GroupVersionKind
	for _, v := range versions {
		result = append(result, definition.GroupVersionKind{
			Group:   group,
			Version: v.Name,
			Kind:    kind,
		})
	}
	return result
}

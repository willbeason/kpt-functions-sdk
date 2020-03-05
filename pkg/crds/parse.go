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

	// Assume the CRD is for a Kubernetes Object.
	// TODO(willbeason): Ensure the fields as required by the Kubernetes API
	//  conventions exist, even if not explicitly defined.
	o := parseObject(*schema)
	o.Meta = definition.Meta{
		Name:        spec.Names.Kind,
		Package:     fmt.Sprintf("%s.%s", spec.Group, version.Name),
		Description: schema.Description,
	}
	o.GroupVersionKinds = gvks(spec.Group, spec.Names.Kind, spec.Versions)
	return o
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

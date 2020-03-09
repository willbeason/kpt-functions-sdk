package crds

import (
	"github.com/willbeason/typegen/pkg/jsonschema"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Parse parses a CRD into a definition.Definition.
func Parse(props *v1.JSONSchemaProps) jsonschema.Type {
	if props == nil {
		return jsonschema.Any{}
	}
}

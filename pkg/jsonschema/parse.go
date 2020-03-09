package jsonschema

import v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Parse parses a CRD into a definition.Definition.
func Parse(name string, props *v1.JSONSchemaProps) Type {
	if props == nil {
		return Any{}
	}

	meta := Meta{
		Name:        name,
		Description: props.Description,
	}



	return nil
}

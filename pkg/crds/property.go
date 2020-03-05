package crds

import (
	"github.com/willbeason/typegen/pkg/definition"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func parseProperty(prop v1.JSONSchemaProps) definition.Property {
	return definition.Property{
		Type:          nil,
		Description:   prop.Description,
		// TODO(willbeason): Set override value.
	}
}

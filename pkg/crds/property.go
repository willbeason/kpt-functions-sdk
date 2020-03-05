package crds

import (
	"github.com/willbeason/typegen/pkg/definition"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func parseProperty(prop v1.JSONSchemaProps) definition.Property {
	return definition.Property{
		Type:          parseType(prop),
		Description:   prop.Description,
		Required:      false,
		// TODO: set override values.
	}
}

func parseType(prop v1.JSONSchemaProps) definition.Type {
	switch prop.Type {
	case definition.NULL, definition.BOOLEAN, definition.INTEGER, definition.NUMBER, definition.STRING:
		return parsePrimitive(prop.Type, prop)
	case definition.ARRAY:
		return parseArray(prop)
	case definition.OBJECT, "":
		return parseObject(prop)
	}
	return nil
}

func parsePrimitive(ts string, p v1.JSONSchemaProps) definition.Primitive {
	return definition.Primitive{
		Type:   ts,
		Format: p.Format,
	}
}

func parseArray(p v1.JSONSchemaProps) definition.Array {
	return definition.Array{
		Items:  parseType(*p.Items.Schema),
	}
}

func parseObject(p v1.JSONSchemaProps) definition.Object {
	def := definition.Object{
		Properties: make(map[string]definition.Property),
		IsKubernetesObject: true,
	}

	for name, property := range p.Properties {
		def.Properties[name] = parseProperty(property)
	}

	for _, name := range p.Required {
		p := def.Properties[name]
		p.Required = true
		def.Properties[name] = p
	}

	return def
}

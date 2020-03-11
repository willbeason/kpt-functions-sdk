package jsonschema

import v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

const (
	AnyType = ""
	BooleanType = "boolean"
	StringType = "string"
	NumberType = "number"
	IntegerType = "integer"
	ArrayType = "array"
	ObjectType = "object"
)

// Parse parses a CRD into a definition.Definition.
func Parse(name string, path []string, props v1.JSONSchemaProps) (Schema, error) {
	path = append([]string{name}, path...)
	meta := Meta{
		Path: path,
		Description: props.Description,
	}

	switch props.Type {
	case AnyType:
		return Any{Meta: meta}, nil
	case BooleanType:
		return Boolean{Meta: meta}, nil
	case StringType:
		return String{Meta: meta}, nil
	case NumberType:
		return parseNumber(meta, props)
	case IntegerType:
		return parseInteger(meta, props)
	case ArrayType:
		return parseArray(meta, props)
	case ObjectType:
		return parseObject(meta, props)
	}

	if props.Ref != nil {
		return Ref{
			Meta: meta,
			Ref: *props.Ref,
		}, nil
	}

	return nil, ParseErrorf(path, "%q must define a schema or %q must be defined", props.Type, "$ref")
}

func parseNumber(meta Meta, props v1.JSONSchemaProps) (Schema, error) {
	format := NumberFormat(props.Format)
	switch format {
	case None, Float, Double, Int32, Int64:
		return Number{Meta: meta, Format: format}, nil
	default:
		return nil, ParseErrorf(meta.Path, "unsupported %q format %q", NumberType, format)
	}
}

func parseInteger(meta Meta, props v1.JSONSchemaProps) (Schema, error) {
	format := NumberFormat(props.Format)
	switch format {
	case None:
		return Number{Meta: meta, Format: Integer}, nil
	case Int32, Int64:
		return Number{Meta: meta, Format: format}, nil
	default:
		return nil, ParseErrorf(meta.Path, "unsupported %q format %q", IntegerType, format)
	}
}

func parseArray(meta Meta, props v1.JSONSchemaProps) (Schema, error) {
	var schema v1.JSONSchemaProps
	switch {
	case props.Items == nil:
		return nil, ParseError(meta.Path, "'items' is required in arrays")
	case props.Items.Schema != nil:
		schema = *props.Items.Schema
	case len(props.Items.JSONSchemas) == 1:
		schema = props.Items.JSONSchemas[0]
	default:
		return nil, ParseErrorf(meta.Path, "'items' must define a single subschema")
	}

	items, err := Parse("items", meta.Path, schema)
	if err != nil {
		return nil, err
	}
	return Array{
		Meta:  meta,
		Items: items,
	}, nil
}

func parseObject(meta Meta, props v1.JSONSchemaProps) (Schema, error) {
	var properties []Schema
	for name, p := range props.Properties {
		prop, err := Parse(name, meta.Path, p)
		if err != nil {
			return nil, err
		}
		properties = append(properties, prop)
	}
	return Object{
		Meta:       meta,
		Properties: properties,
	}, nil
}

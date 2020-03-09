package jsonschema

// Type is one of five non-null primitive types defined in the JSON Schema
// specification.
//
// https://tools.ietf.org/html/draft-handrews-json-schema-02#section-4.2.1
//
// Null and mixed types are not permitted in the Open API schema, so any field
// will be exactly one of the five non-null primitive types.
//
// https://swagger.io/docs/specification/data-models/data-types/
type Type interface {
	// Children is the Objects defined by this Type.
	Children() []Type
}

// Any is a schema that does not specify "type".
type Any struct {
	Meta
}

func (a Any) Children() []Type {
	return nil
}

// Boolean is Type that may take the value "true" or "false".
type Boolean struct {
	Meta
}

// Children implements Type.
func (b Boolean) Children() []Type {
	return nil
}

// NumberFormat represents one of the six recognized formats for number types
// in the OpenAPI Schema.
//
// https://swagger.io/docs/specification/data-models/data-types/#numbers
type NumberFormat int

const (
	Float = iota
	Double
	Integer
	Int32
	Int64
)

// Number is a numeric Type, specified by the Format.
type Number struct {
	Meta

	// Format specifies how the number should be represented.
	// If unset, defaults to Double.
	Format *NumberFormat

	// TODO(willbeason): Support other specified Number validation per
	//  https://swagger.io/docs/specification/data-models/data-types/#numbers
}

// Children implements Type.
func (n Number) Children() []Type {
	return nil
}

// String is Type consisting of a sequence of Unicode code points.
type String struct {
	Meta

	// TODO(willbeason): Support string format validation per
	//  https://swagger.io/docs/specification/data-models/data-types/#string
}

// Array is a Type consisting of a sequence of primitives.
//
// https://swagger.io/docs/specification/data-models/data-types/#array
type Array struct {
	Meta

	// Items is the type of all members contained in Arrays of this Type.
	Items Type

	// TODO(willbeason): Support array validation per specification.
}

// Children implements Type.
func (a Array) Children() []Type {
	return []Type{a.Items}
}

// Object is a collection of property/value pairs.
type Object struct {
	Meta

	Properties []Type
}

// Children implements Type.
func (o Object) Children() []Type {
	return o.Properties
}

package jsonschema

// Schema is one of five non-null primitive types defined in the JSON Schema
// specification, the Any Schema, or a Reference to another Schema.
//
// https://tools.ietf.org/html/draft-handrews-json-schema-02#section-4.2.1
//
// Null and mixed types are not permitted in the Open API schema.
//
// https://swagger.io/docs/specification/data-models/data-types/
type Schema interface {
	// Subschemas are the in-line Types defined by this Schema.
	Subschemas() []Schema

	// Name is the unqualified name of the schema.
	Name() string

	Comment() string
}

// Any is a schema that does not specify "type".
//
// https://swagger.io/docs/specification/data-models/data-types/#any
type Any struct {
	Meta
}

// Subschemas implements Schema.
func (a Any) Subschemas() []Schema {
	return nil
}

// Boolean is Schema whose instances may take the value "true" or "false".
type Boolean struct {
	Meta
}

// Subschemas implements Schema.
func (b Boolean) Subschemas() []Schema {
	return nil
}

// NumberFormat represents one of the six recognized formats for number types
// in the OpenAPI Schema.
//
// https://swagger.io/docs/specification/data-models/data-types/#numbers
type NumberFormat string

const (
	None NumberFormat = ""
	Float = "float"
	Double = "double"
	Integer = "integer"
	Int32 = "int32"
	Int64 = "int64"
)

// Number is a numeric Schema, specified by the Format.
type Number struct {
	Meta

	// Format specifies how the number should be represented.
	Format NumberFormat

	// TODO(willbeason): Support number validation per
	//  https://swagger.io/docs/specification/data-models/data-types/#numbers
}

// Subschemas implements Schema.
func (n Number) Subschemas() []Schema {
	return nil
}

// String is Schema defining of a sequence of Unicode code points.
type String struct {
	Meta

	// TODO(willbeason): Support string format validation per
	//  https://swagger.io/docs/specification/data-models/data-types/#string
}

func (s String) Subschemas() []Schema {
	return nil
}

// Array is a Schema defining of a sequence of primitives.
//
// https://swagger.io/docs/specification/data-models/data-types/#array
type Array struct {
	Meta

	// Items is the type of all members contained in Arrays of this Schema.
	Items Schema

	// TODO(willbeason): Support array validation per specification.
}

// Subschemas implements Schema.
func (a Array) Subschemas() []Schema {
	return []Schema{a.Items}
}

// Object is a collection of property/value pairs.
type Object struct {
	Meta

	Properties []Schema
}

// Subschemas implements Schema.
func (o Object) Subschemas() []Schema {
	return o.Properties
}

// Ref is a reference to another Schema.
type Ref struct {
	Meta

	Ref string
}

// Subschemas implements Schema.
func (r Ref) Subschemas() []Schema {
	return nil
}

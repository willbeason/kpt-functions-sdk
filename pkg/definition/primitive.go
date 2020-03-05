package definition

const (
	BOOLEAN = "boolean"
	INTEGER = "integer"
	NUMBER  = "number"
	STRING  = "string"
	ARRAY   = "array"
	OBJECT  = "object"
)

// KnownPrimitives are the set of base "type"/"format" combinations defined here:
// https://swagger.io/docs/specification/data-models/data-types/
//
// Any combination is valid; it is up to the language implementer to decide what to do with an unknown combination and
// set reasonable defaults when "format" is unexpected or "type" is unknown. Generally it is safest to default to
// "string".
var KnownPrimitives = struct {
	Boolean Primitive
	Integer Primitive
	Long    Primitive
	Float   Primitive
	Double  Primitive
	String  Primitive

	// Byte is base64-encoded characters.
	Byte Primitive

	// Binary is a sequence of octets.
	Binary Primitive

	// Date is an RFC3339 date.
	Date Primitive

	// DateTime is an RFC3339 date/time.
	DateTime Primitive

	// Password should never appear.
	Password Primitive
}{
	Boolean:  Primitive{Type: BOOLEAN},
	Integer:  Primitive{Type: INTEGER, Format: "int32"},
	Long:     Primitive{Type: INTEGER, Format: "int64"},
	Float:    Primitive{Type: NUMBER, Format: "float"},
	Double:   Primitive{Type: NUMBER, Format: "double"},
	String:   Primitive{Type: STRING},
	Byte:     Primitive{Type: STRING, Format: "byte"},
	Binary:   Primitive{Type: STRING, Format: "binary"},
	Date:     Primitive{Type: STRING, Format: "date"},
	DateTime: Primitive{Type: STRING, Format: "date-time"},
	Password: Primitive{Type: STRING, Format: "password"},
}

// Primitive describes a Primitive type and its Format.
//
// These usually correspond to primitive types in languages like "int32".
type Primitive struct {
	Type   string // "type"
	Format string // "Format"
}

var _ Type = Primitive{}

// Imports implements Type.
func (Primitive) Imports() []Ref {
	// Primitives never require any imports to use.
	return nil
}

// NestedTypes implements Type.
func (Primitive) NestedTypes() []Object {
	// Primitives never have nested types.
	return nil
}
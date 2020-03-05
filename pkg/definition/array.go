package definition

// Array is an array of any Type.
type Array struct {
	// Items is the type of elements in the Array. Defined in the "Items" field.
	Items Type

	// Nested is the set of types nested in this Array. Guaranteed to be either empty or contain exactly one element.
	Nested []Object
}

var _ Type = Array{}

// Imports implements Type.
func (a Array) Imports() []Ref {
	// An Array's imports may be either the Item it is an array of or the imports required for its nested field.
	return a.Items.Imports()
}

// NestedTypes implements Type.
func (a Array) NestedTypes() []Object {
	// Arrays may define their "items" type inline, creating a nested field.
	return a.Nested
}
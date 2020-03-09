package definition

// Type represents a swagger.json type in a field declaration.
type Type interface {
	// Imports is the set of imports required to use this Type.
	Imports() []Ref

	// NestedTypes returns the type definitions nested in this type.
	NestedTypes() []Object
}

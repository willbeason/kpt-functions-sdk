package definition

// Map is a map from strings to any Type.
//
// Corresponds to swagger's Free-Form Object with "additionalProperties" but no "properties" defined.
type Map struct {
	// Values is the type of Values in the map. Defined in the "additionalProperties" field.
	Values Type
}

var _ Type = Map{}

// Imports implements Type.
func (m Map) Imports() []Ref {
	return m.Values.Imports()
}

// NestedTypes implements Type.
func (Map) NestedTypes() []Object {
	// Will change once we support maps of nested type.
	return nil
}

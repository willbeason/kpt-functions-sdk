package definition

import "sort"

// Object corresponds to structured Objects with "properties" but no "additionalProperties".
type Object struct {
	Meta

	// Properties is the set of properties this Object declares.
	Properties map[string]Property

	// NestedTypes are all types which only exist inside this Object's definition.
	NestedTypes []Object

	// IsKubernetesObject is true if Object represents a KubernetesObject type.
	IsKubernetesObject bool

	// GroupVersionKinds holds the set of declared GroupVersionKinds this type may validly have.
	// Must be nonempty to be a KubernetesObject. A Definition may declare many GroupVersionKinds.
	GroupVersionKinds []GroupVersionKind
}

// NamedProperty is a Property and its name.
type NamedProperty struct {
	Name string
	Property
}

// NamedProperties returns the Properties of the Object, sorted by their name.
func (o Object) NamedProperties() []NamedProperty {
	var namedProperties []NamedProperty
	for name, property := range o.Properties {
		namedProperties = append(namedProperties, NamedProperty{Name: name, Property: property})
	}
	sort.Slice(namedProperties, func(i, j int) bool {
		return namedProperties[i].Name < namedProperties[j].Name
	})
	return namedProperties
}

// GroupVersionKind returns the first listed GroupVersionKind for the type, if there is one.
func (o Object) GroupVersionKind() *GroupVersionKind {
	// TODO(b/142004166): Handle the case where a type declares multiple GroupVersionKinds.
	if len(o.GroupVersionKinds) > 0 {
		return &o.GroupVersionKinds[0]
	}
	return nil
}

// HasRequiredFields returns true if the Object contains any required fields.
func (o Object) HasRequiredFields() bool {
	for _, property := range o.Properties {
		if property.Required {
			return true
		}
	}
	return false
}

// Imports implements Definition.
func (o Object) Imports() []Ref {
	var result []Ref
	for _, p := range o.Properties {
		result = append(result, p.Type.Imports()...)
	}
	for _, n := range o.NestedTypes {
		result = append(result, n.Imports()...)
	}
	return result
}
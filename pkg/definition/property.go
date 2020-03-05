package definition

type Property struct {
	Type          Type
	Description   string
	Required      bool
	OverrideValue string
}

func (p Property) WithRequired() Property {
	p.Required = true
	return p
}

func (p Property) WithType(t Type) Property {
	p.Type = t
	return p
}
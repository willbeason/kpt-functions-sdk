package jsonschema

type Meta struct {
	// Path is the path to this Schema.
	Path []string

	// Description is an annotation containing a human-readable textual comment describing this Schema.
	Description string
}

func (m Meta) Name() string {
	return m.Path[0]
}

func (m Meta) Comment() string {
	return m.Description
}

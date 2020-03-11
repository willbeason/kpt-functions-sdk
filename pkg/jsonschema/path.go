package jsonschema

import "strings"

type Path []string

func (p Path) String() string {
	return strings.Join(p, ".")
}

package ts

import (
	"fmt"
	"github.com/willbeason/typegen/pkg/jsonschema"
	"strings"
)

func indent(strs []string) []string {
	var result []string
	for _, s := range strs {
		result = append(result, "  "+s)
	}
	return result
}

func Print(schema jsonschema.Schema) string {
	var result []string

	result = append(result, printComment(schema.Comment())...)
	switch s := schema.(type) {
	case jsonschema.Any:
		result = append(result, printAny(s))
	case jsonschema.Object:
		result = append(result, printObject(s)...)
	default:
		panic(fmt.Sprintf("unsupported type %T", s))
	}

	return strings.Join(result, "\n")
}

func print(schema jsonschema.Schema) string {

}

func printComment(comment string) []string {
	var result []string
	lines := strings.Split(comment, "\n")
	for _, line := range lines {
		result = append(result, "// " + line)
	}
	return result
}


package ts

import (
	"fmt"
	"github.com/willbeason/typegen/pkg/jsonschema"
)

func printObject(o jsonschema.Object) []string {
	var result []string
	result = append(result, fmt.Sprintf("export class %s implements KubernetesObject {\n", o.Name()))
	result = append(result, "}")

	return result
}
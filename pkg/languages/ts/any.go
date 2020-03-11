package ts

import (
	"fmt"
	"github.com/willbeason/typegen/pkg/jsonschema"
)

func printAny(any jsonschema.Any) string {
	return fmt.Sprintf("export type %s = object", any.Name())
}

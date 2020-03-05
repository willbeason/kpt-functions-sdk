package swagger

import (
	"fmt"
	"github.com/willbeason/typegen/pkg/definition"
	"github.com/willbeason/typegen/pkg/maps"
	"strings"
)

func parseRef(r map[string]interface{}) definition.Ref {
	ref := maps.GetRequiredString("$ref", r)

	if !strings.HasPrefix(ref, "#/definitions/") {
		panic(fmt.Sprintf("invalid $ref, must begin with '#/definitions/': %s", r))
	}
	i := strings.LastIndex(ref, ".")

	return definition.Ref{
		Package: ref[14:i],
		Name:    ref[i+1:],
	}
}
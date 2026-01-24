package templUtil

import (
	"github.com/a-h/templ"
	"github.com/gungun974/gocva"
	"maps"
)

func ToJSONString(object any) string {
	JSONStr, err := templ.JSONString(object)
	if err != nil {
		JSONStr = "{}"
	}

	return JSONStr
}

func MergeVariant(newVariant gocva.Variant, oldVariant gocva.Variant) gocva.Variant {
	if newVariant == nil {
		return oldVariant
	}
	return newVariant
}

func MergeAttrs(newAttrs templ.Attributes, oldAttrs templ.Attributes) templ.Attributes {
	if newAttrs == nil {
		return oldAttrs
	}

	maps.Copy(oldAttrs, newAttrs)
	return oldAttrs
}

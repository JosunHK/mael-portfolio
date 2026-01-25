package templUtil

import (
	"database/sql"
	"github.com/a-h/templ"
	"github.com/gungun974/gocva"
	"maps"
	"strconv"
)

func ToJSONString(object any) string {
	JSONStr, err := templ.JSONString(object)
	if err != nil {
		JSONStr = "{}"
	}

	return JSONStr
}

func MergeVariant(oldVariant gocva.Variant, newVariant gocva.Variant) gocva.Variant {
	if newVariant == nil {
		return oldVariant
	}
	return newVariant
}

func MergeAttrs(oldAttrs templ.Attributes, newAttrs templ.Attributes) templ.Attributes {
	if newAttrs == nil {
		return oldAttrs
	}

	maps.Copy(oldAttrs, newAttrs)
	return oldAttrs
}

// general merger
func MergeAttr(oldAttr string, newAttr string) string {
	if newAttr == "" {
		return oldAttr
	}

	return newAttr
}

func ExtractInt32(data sql.NullInt32) string {
	if !data.Valid {
		return ""
	}

	return strconv.Itoa(int(data.Int32))
}

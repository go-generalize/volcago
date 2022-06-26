package generator

import (
	"go/types"
	"reflect"
	"strings"

	eptypes "github.com/go-generalize/go-easyparser/types"
)

const (
	typeString    = "string"
	typeInt       = "int"
	typeInt64     = "int64"
	typeFloat64   = "float64"
	typeBool      = "bool"
	_             = "interface{}"
	_             = "time.Time"
	_             = "*time.Time"
	typeLatLng    = "latlng.LatLng"
	typeReference = "firestore.DocumentRef"
	typeMap       = "map[string]"
	_             = "map[string]string"
	_             = "map[string]int"
	_             = "map[string]int64"
	_             = "map[string]float64"
	typeBoolMap   = "map[string]bool"
	_             = "map[string]interface{}"
)

func getGoTypeFromEPTypes(t eptypes.Type) string {
	switch t := t.(type) {
	case *eptypes.String:
		return "string"
	case *eptypes.Number:
		switch t.RawType {
		case types.Int:
			return "int"
		case types.Int8:
			return "int8"
		case types.Int16:
			return "int16"
		case types.Int32:
			return "int32"
		case types.Int64:
			return "int64"
		case types.Uint:
			return "uint"
		case types.Uint8:
			return "uint8"
		case types.Uint16:
			return "uint16"
		case types.Uint32:
			return "uint32"
		case types.Uint64:
			return "uint64"
		case types.Uintptr:
			return "uintptr"
		case types.Float32:
			return "float32"
		case types.Float64:
			return "float64"
		}
	case *eptypes.Boolean:
		return "bool"
	case *eptypes.Nullable:
		r := getGoTypeFromEPTypes(t.Inner)

		if r == "" {
			return ""
		}

		if strings.HasPrefix(r, "[]") || strings.HasPrefix(r, "map[") {
			return r
		}

		return "*" + r
	case *eptypes.Array:
		return "[]" + getGoTypeFromEPTypes(t.Inner)
	case *eptypes.Date:
		return "time.Time"
	case *eptypes.Object:
		names := strings.Split(t.Name, ".")
		return names[len(names)-1]
	case *eptypes.Map:
		return "map[" + getGoTypeFromEPTypes(t.Key) + "]" + getGoTypeFromEPTypes(t.Value)
	case *eptypes.Any:
		return "interface{}"
	case *documentRef:
		return typeReference
	case *latLng:
		return typeLatLng
	}

	panic("unsupported: " + reflect.TypeOf(t).String())
}

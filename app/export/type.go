package export

import (
	"fmt"

	types2 "github.com/kyleu/admini/app/schema/types"
	"github.com/kyleu/admini/app/util"
)

// nolint
func typeString(typ types2.Type, f *Format, ctx string) (string, []util.Pkg) {
	switch t := typ.(type) {
	case *types2.Wrapped:
		return typeString(t.T, f, ctx)
	case *types2.Unknown:
		return fmt.Sprintf("string /* %s */", t.X), nil
	case *types2.Error:
		return "string /* ERROR: " + t.Message + " */", nil

	case *types2.Bool:
		return "bool", nil
	case *types2.Int:
		return "int", nil
	case *types2.Enum:
		return "string /* " + t.Ref + " */", nil
	case *types2.JSON:
		switch ctx {
		case "dto":
			return "json.RawMessage", []util.Pkg{{"encoding/json"}}
		default:
			if t.IsObject {
				return "util.ValueMap", nil
			}
			if t.IsArray {
				return "[]interface{}", nil
			}
			return "interface{}", nil
		}
	case *types2.List:
		ts, p := typeString(t.V, f, ctx)
		return "[]" + ts, p
	case *types2.Map:
		kts, kp := typeString(t.K, f, ctx)
		vts, vp := typeString(t.V, f, ctx)
		return fmt.Sprintf("map[%s]%s", kts, vts), append(kp, vp...)
	case *types2.Float:
		return "float", nil
	case *types2.Option:
		if ctx == "dto" {
			switch t.V.T.(type) {
			case *types2.Bool:
				return "sql.NullBool", []util.Pkg{{"database/sql"}}
			case *types2.String:
				return "sql.NullString", []util.Pkg{{"database/sql"}}
			}
		}
		ts, p := typeString(t.V, f, ctx)
		return "*" + ts, p
	case *types2.String:
		return "string", nil
	case *types2.Timestamp, *types2.TimestampZoned:
		return "time.Time", nil
	case *types2.UUID:
		return "uuid.UUID", nil
	default:
		return "string /* " + t.String() + " */", nil
	}
}

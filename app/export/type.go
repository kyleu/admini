package export

import (
	"fmt"

	"github.com/kyleu/admini/app/types"
	"github.com/kyleu/admini/app/util"
)

// nolint
func typeString(typ types.Type, f *Format, ctx string) (string, []util.Pkg) {
	switch t := typ.(type) {
	case *types.Wrapped:
		return typeString(t.T, f, ctx)
	case *types.Unknown:
		return fmt.Sprintf("string /* %s */", t.X), nil
	case *types.Error:
		return "string /* ERROR: " + t.Message + " */", nil

	case *types.Bool:
		return "bool", nil
	case *types.Int:
		return "int", nil
	case *types.Enum:
		return "string /* " + t.Ref + " */", nil
	case *types.JSON:
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
	case *types.List:
		ts, p := typeString(t.V, f, ctx)
		return "[]" + ts, p
	case *types.Map:
		kts, kp := typeString(t.K, f, ctx)
		vts, vp := typeString(t.V, f, ctx)
		return fmt.Sprintf("map[%s]%s", kts, vts), append(kp, vp...)
	case *types.Float:
		return "float", nil
	case *types.Option:
		if ctx == "dto" {
			switch t.V.T.(type) {
			case *types.Bool:
				return "sql.NullBool", []util.Pkg{{"database/sql"}}
			case *types.String:
				return "sql.NullString", []util.Pkg{{"database/sql"}}
			}
		}
		ts, p := typeString(t.V, f, ctx)
		return "*" + ts, p
	case *types.String:
		return "string", nil
	case *types.Timestamp, *types.TimestampZoned:
		return "time.Time", nil
	case *types.UUID:
		return "uuid.UUID", nil
	default:
		return "string /* " + t.String() + " */", nil
	}
}

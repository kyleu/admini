package export

import (
	"fmt"

	"github.com/kyleu/admini/app/schema/schematypes"
	"github.com/kyleu/admini/app/util"
)

func typeString(typ schematypes.Type, f *Format, ctx string) (string, []util.Pkg) {
	switch t := typ.(type) {
	case *schematypes.Wrapped:
		return typeString(t.T, f, ctx)
	case *schematypes.Unknown:
		return fmt.Sprintf("string /* %s */", t.X), nil
	case *schematypes.Error:
		return "string /* ERROR: " + t.Message + " */", nil

	case *schematypes.Bool:
		return "bool", nil
	case *schematypes.Int:
		return "int", nil
	case *schematypes.JSON:
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
	case *schematypes.List:
		ts, p := typeString(t.V, f, ctx)
		return "[]" + ts, p
	case *schematypes.Map:
		kts, kp := typeString(t.K, f, ctx)
		vts, vp := typeString(t.V, f, ctx)
		return fmt.Sprintf("map[%s]%s", kts, vts), append(kp, vp...)
	case *schematypes.Float:
		return "float", nil
	case *schematypes.Option:
		if ctx == "dto" {
			switch t.V.T.(type) {
			case *schematypes.Bool:
				return "sql.NullBool", []util.Pkg{{"database/sql"}}
			case *schematypes.String:
				return "sql.NullString", []util.Pkg{{"database/sql"}}
			}
		}
		ts, p := typeString(t.V, f, ctx)
		return "*" + ts, p
	case *schematypes.String:
		return "string", nil
	case *schematypes.Timestamp, *schematypes.TimestampZoned:
		return "time.Time", nil
	case *schematypes.UUID:
		return "uuid.UUID", nil
	default:
		return "string /* " + t.String() + " */", nil
	}
}

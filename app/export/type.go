package export

import (
	"github.com/kyleu/admini/app/schema/schematypes"
	"github.com/kyleu/admini/app/util"
)

func typeString(typ schematypes.Type, f *Format) (string, []util.Pkg) {
	switch t := typ.(type) {
	case *schematypes.Wrapped:
		return typeString(t.T, f)
	case *schematypes.Unknown:
		return "string /* " + t.X + " */", nil
	case *schematypes.Error:
		return "string /* ERROR: " + t.Message + " */", nil

	case *schematypes.Bool:
		return "bool", nil
	case *schematypes.Int:
		return "int", nil
	case *schematypes.JSON:
		return "json.RawMessage", nil
	case *schematypes.List:
		ts, p := typeString(t.T, f)
		return "[]" + ts, p
	case *schematypes.Map:
		kts, kp := typeString(t.K, f)
		vts, vp := typeString(t.V, f)
		return "map[" + kts + "]" + vts, append(kp, vp...)
	case *schematypes.Option:
		ts, p := typeString(t.T, f)
		return "*" + ts, p
	case *schematypes.String:
		return "string", nil
	case *schematypes.Timestamp:
		return "time.Time", nil
	case *schematypes.UUID:
		return "uuid.UUID", nil
	default:
		return "string /* " + t.String() + " */", nil
	}
}
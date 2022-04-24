package schema

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/lib/types"
	"admini.dev/admini/app/util"
)

func (s *Schema) Hack(logger *zap.SugaredLogger) (string, error) {
	ret := make([]util.ValueMap, 0, len(s.Models))
	for _, m := range s.Models {
		if m.Type != model.TypeEnum {
			ret = append(ret, hackModel(m, logger))
		}
	}
	return util.ToJSON(ret), nil
}

func hackModel(m *model.Model, logger *zap.SugaredLogger) util.ValueMap {
	cols := make([]util.ValueMap, 0, len(m.Fields))
	for _, f := range m.Fields {
		cols = append(cols, hackField(m, f, logger))
	}
	return util.ValueMap{
		"columns":     cols,
		"description": m.Description(),
		"icon":        "star",
		"name":        m.Key,
		"ordering":    []any{},
		"package":     strings.ToLower(m.Key),
		"search":      []any{},
	}
}

func hackField(m *model.Model, f *field.Field, logger *zap.SugaredLogger) util.ValueMap {
	ret := util.ValueMap{
		"name":     f.Key,
		"pk":       m.IsPK(f.Key, logger),
		"nullable": f.Type.IsOption(),
	}
	ret["type"] = typeFor(f.Type)
	if f.Default != nil {
		ret["sqlDefault"] = fmt.Sprint(f.Default)
	}
	return ret
}

func typeFor(t types.Type) string {
	if w, ok := t.(*types.Wrapped); ok {
		return typeFor(w.T)
	}
	if o, ok := t.(*types.Option); ok {
		return typeFor(o.V)
	}
	switch t.Key() {
	case "timestampZoned":
		return "timestamp"
	default:
		return t.Key()
	}
}

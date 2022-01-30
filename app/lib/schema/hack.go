package schema

import (
	"fmt"
	"strings"

	types2 "github.com/kyleu/admini/app/lib/types"
	"go.uber.org/zap"

	"github.com/kyleu/admini/app/lib/schema/field"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/kyleu/admini/app/util"
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
		"ordering":    []interface{}{},
		"package":     strings.ToLower(m.Key),
		"search":      []interface{}{},
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

func typeFor(t types2.Type) string {
	if w, ok := t.(*types2.Wrapped); ok {
		return typeFor(w.T)
	}
	if o, ok := t.(*types2.Option); ok {
		return typeFor(o.V)
	}
	switch t.Key() {
	case "timestampZoned":
		return "timestamp"
	default:
		return t.Key()
	}
}

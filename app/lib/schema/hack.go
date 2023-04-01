package schema

import (
	"fmt"
	"os"
	"strings"

	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/lib/types"
	"admini.dev/admini/app/util"
)

func (s *Schema) Hack(logger util.Logger) (string, error) {
	ret := make([]util.ValueMap, 0, len(s.Models))
	for _, m := range s.Models {
		if m.Type != model.TypeEnum {
			md := hackModel(m, logger)
			err := os.WriteFile("./tmp/"+m.Key+".json", util.ToJSONBytes(md, true), 0644)
			if err != nil {
				return "", err
			}
			ret = append(ret, md)
		}
	}
	return util.ToJSON(ret), nil
}

func (s *Schema) HackSvc(logger util.Logger) (string, error) {
	var ret string
	for _, m := range s.Models {
		if m.Type != model.TypeEnum {
			ret += fmt.Sprintf("%s *%s.Service\n", m.Key, strings.ToLower(m.Key))
		}
	}
	ret += "\n\n\n"
	for _, m := range s.Models {
		if m.Type != model.TypeEnum {
			ret += fmt.Sprintf("%s: %s.NewService(st.DB),\n", m.Key, strings.ToLower(m.Key))
		}
	}
	return ret, nil
}

func hackModel(m *model.Model, logger util.Logger) util.ValueMap {
	cols := make([]util.ValueMap, 0, len(m.Fields))
	for _, f := range m.Fields {
		cols = append(cols, hackField(m, f, logger))
	}
	rels := make([]util.ValueMap, 0, len(m.Relationships))
	for _, r := range m.Relationships {
		rels = append(rels, hackRel(r))
	}
	return util.ValueMap{
		"columns":     cols,
		"description": m.Description(),
		"icon":        "star",
		"name":        m.Key,
		"ordering":    []any{},
		"package":     strings.ToLower(m.Key),
		"relations":   rels,
		"search":      []any{},
	}
}

func hackField(m *model.Model, f *field.Field, logger util.Logger) util.ValueMap {
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

func hackRel(r *model.Relationship) util.ValueMap {
	ret := util.ValueMap{
		"name":  r.Key,
		"src":   r.SourceFields,
		"table": r.TargetModel,
		"tgt":   r.TargetFields,
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
	case "timestampZoned", "time":
		return "timestamp"
	default:
		return t.Key()
	}
}

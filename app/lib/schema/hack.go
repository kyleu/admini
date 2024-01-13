package schema

import (
	"fmt"
	"os"
	"strings"

	"admini.dev/admini/app/lib/filesystem"
	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/lib/types"
	"admini.dev/admini/app/util"
)

func (s *Schema) Hack(logger util.Logger) (string, error) {
	ret := make([]*util.OrderedMap[any], 0, len(s.Models))
	for _, m := range s.Models {
		if m.Type != model.TypeEnum {
			md := hackModel(m, logger)
			err := os.WriteFile("./tmp/"+m.Key+".json", util.ToJSONBytes(md, true), os.FileMode(filesystem.DefaultMode))
			if err != nil {
				return "", err
			}
			ret = append(ret, md)
		}
	}
	return util.ToJSON(ret), nil
}

func (s *Schema) HackSvc(_ util.Logger) (string, error) {
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

func hackModel(m *model.Model, logger util.Logger) *util.OrderedMap[any] {
	cols := make([]*util.OrderedMap[any], 0, len(m.Fields))
	for _, f := range m.Fields {
		cols = append(cols, hackField(m, f, logger))
	}
	rels := make([]util.ValueMap, 0, len(m.Relationships))
	for _, r := range m.Relationships {
		rels = append(rels, hackRel(r))
	}
	ret := util.NewOrderedMap[any](false, 8)
	ret.Append("name", m.Key)
	ret.Append("package", strings.ToLower(m.Key))
	ret.Append("description", m.Description())
	ret.Append("icon", "star")
	ret.Append("ordering", []any{})
	ret.Append("search", []any{})
	ret.Append("columns", cols)
	ret.Append("relations", rels)
	return ret
}

func hackField(m *model.Model, f *field.Field, logger util.Logger) *util.OrderedMap[any] {
	ret := util.NewOrderedMap[any](false, 3)
	ret.Append("name", f.Key)
	ret.Append("type", typeFor(f.Type))
	ret.Append("nullable", f.Type.IsOption())
	ret.Append("pk", m.IsPK(f.Key, logger))
	if f.Default != nil {
		ret.Append("sqlDefault", fmt.Sprint(f.Default))
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

package schema

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
)

const (
	KeyTitle  = "title"
	KeyPlural = "plural"
)

type Override struct {
	Type string   `json:"type,omitempty"`
	Path util.Pkg `json:"path,omitempty"`
	Prop string   `json:"prop,omitempty"`
	Val  any      `json:"val,omitempty"`
}

func NewOverride(typ string, path util.Pkg, prop string, val any) *Override {
	return &Override{Type: typ, Path: path, Prop: prop, Val: val}
}

func (o Override) String() string {
	return fmt.Sprintf("%s[%s].%s = %v", o.Type, strings.Join(o.Path, "/"), o.Prop, o.Val)
}

type Overrides []*Override

func (o Overrides) Purge(path util.Pkg) Overrides {
	return lo.Filter(o, func(x *Override, _ int) bool {
		return !(x.Path.StartsWith(path))
	})
}

func (o Overrides) Sort() {
	slices.SortFunc(o, func(l *Override, r *Override) int {
		if !l.Path.Equals(r.Path) {
			for idx, p := range l.Path {
				if idx >= len(r.Path) {
					return -1
				}
				if p != r.Path[idx] {
					return cmp.Compare(p, r.Path[idx])
				}
			}
		}
		if len(l.Path) != len(r.Path) {
			return -1
		}
		if l.Prop != r.Prop {
			return cmp.Compare(l.Prop, r.Prop)
		}
		if l.Type != r.Type {
			return cmp.Compare(l.Type, r.Type)
		}
		return 0
	})
}

func (o *Override) ApplyTo(s *Schema) error {
	switch o.Type {
	case "model":
		m := s.Models.Get(o.Path.Drop(1), o.Path.Last())
		if m == nil {
			return errors.Errorf("unable to find model at path [%s]", strings.Join(o.Path, "/"))
		}
		return applyModelProperty(m, o.Prop, o.Val)
	case "field":
		mPath := o.Path.Drop(1)
		fKey := o.Path.Last()
		m := s.Models.Get(mPath.Drop(1), mPath.Last())
		if m == nil {
			return errors.Errorf("unable to find model at path [%s]", strings.Join(o.Path, "/"))
		}
		_, f := m.Fields.Get(fKey)
		return applyFieldProperty(m, f, o.Prop, o.Val)
	default:
		return errors.Errorf("unhandled override type [%s]", o.Type)
	}
}

func applyModelProperty(m *model.Model, prop string, val any) error {
	switch prop {
	case KeyTitle:
		m.Title = fmt.Sprintf("%v", val)
	case KeyPlural:
		m.Plural = fmt.Sprintf("%v", val)
	default:
		return errors.Errorf("unhandled model property [%s]", prop)
	}
	return nil
}

func applyFieldProperty(_ *model.Model, f *field.Field, prop string, val any) error {
	switch prop {
	case KeyTitle:
		f.Title = fmt.Sprintf("%v", val)
	case KeyPlural:
		f.Plural = fmt.Sprintf("%v", val)
	default:
		return errors.Errorf("unhandled field property [%s]", prop)
	}
	return nil
}

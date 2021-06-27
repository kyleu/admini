package schema

import (
	"fmt"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"strings"
)

type Override struct {
	Type string      `json:"type,omitempty"`
	Path util.Pkg    `json:"path,omitempty"`
	Prop string      `json:"prop,omitempty"`
	Val  interface{} `json:"val,omitempty"`
}

func (o Override) String() string {
	return fmt.Sprintf("%s[%s].%s = %v", o.Type, strings.Join(o.Path, "/"), o.Prop, o.Val)
}

type Overrides []*Override

func (o *Override) ApplyTo(s *Schema) error {
	switch o.Type {
	case "model":
		m := s.Models.Get(o.Path.Drop(1), o.Path.Last())
		if m == nil {
			return errors.Errorf("unable to find model at path [%s]", strings.Join(o.Path, "/"))
		}
		return applyModelProperty(m, o.Prop, o.Val)
	default:
		return errors.Errorf("unhandled override type [%s]", o.Type)
	}
}

func applyModelProperty(m *model.Model, prop string, val interface{}) error {
	switch prop {
	case "title":
		m.Title = fmt.Sprintf("%v", val)
	case "plural":
		m.Plural = fmt.Sprintf("%v", val)
	default:
		return errors.Errorf("unhandled model property [%s]", prop)
	}
	return nil
}

func (s *Schema) CreateReferences() error {
	for _, src := range s.Models {
		if src.References != nil {
			return errors.New("double call of CreateReferences")
		}
		for _, tgt := range s.Models {
			for _, rel := range tgt.Relationships {
				if rel.TargetPkg.Equals(src.Pkg) && rel.TargetModel == src.Key {
					err := src.AddReference(model.ReferenceFromRelation(rel, tgt))
					if err != nil {
						return errors.Wrapf(err, "unable to add reference to [%s]", src.String())
					}
				}
			}
		}
	}
	return nil
}

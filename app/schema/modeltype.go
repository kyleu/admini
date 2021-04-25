package schema

import (
	"encoding/json"

	"github.com/kyleu/admini/app/util"
)

type ModelType struct {
	Key    string
	Title  string
	Plural string
}

var (
	ModelTypeEnum         = ModelType{Key: "enum", Title: "Enum", Plural: "Enums"}
	ModelTypeSequence     = ModelType{Key: "sequence", Title: "Sequence", Plural: "Sequences"}
	ModelTypeStruct       = ModelType{Key: "struct", Title: "Struct", Plural: "Structs"}
	ModelTypeInterface    = ModelType{Key: "interface", Title: "Interface", Plural: "Interfaces"}
	ModelTypeUnion        = ModelType{Key: "union", Title: "Union", Plural: "Unions"}
	ModelTypeIntersection = ModelType{Key: "intersection", Title: "Intersection", Plural: "Intersections"}
	ModelTypeUnknown      = ModelType{Key: "unknown", Title: "Unknown", Plural: "Unknowns"}
)

var AllModelTypes = []ModelType{
	ModelTypeEnum, ModelTypeSequence, ModelTypeStruct,
	ModelTypeInterface, ModelTypeUnion, ModelTypeIntersection,
}

func modelTypeFromString(s string) ModelType {
	for _, t := range AllModelTypes {
		if t.Key == s {
			return t
		}
	}
	util.LogWarn("unhandled model type [" + s + "]")
	return ModelTypeUnknown
}

func (t *ModelType) String() string {
	return t.Key
}

func (t *ModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *ModelType) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = modelTypeFromString(s)
	return nil
}

type ModelPackage struct {
	Key           string        `json:"key"`
	Pkg           util.Pkg      `json:"pkg"`
	ChildModels   Models        `json:"childModels,omitempty"`
	ChildPackages ModelPackages `json:"childPackages,omitempty"`
}

func (m *ModelPackage) Path() string {
	return "/" + m.Pkg.ToPath(m.Key)
}

func (p *ModelPackage) GetPkg(key string, createIfMissing bool) *ModelPackage {
	for _, x := range p.ChildPackages {
		if x.Key == key {
			return x
		}
	}
	if createIfMissing {
		x := &ModelPackage{Key: key, Pkg: append(p.Pkg, p.Key)}
		p.ChildPackages = append(p.ChildPackages, x)
		return x
	}
	return nil
}

func (p *ModelPackage) GetModel(key string) *Model {
	for _, x := range p.ChildModels {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func (p *ModelPackage) Add(pkg util.Pkg, m *Model) {
	if len(pkg) == 0 {
		p.ChildModels = append(p.ChildModels, m)
	} else {
		x := p.GetPkg(pkg[0], true)
		x.Add(pkg[1:], m)
	}
}

func (p *ModelPackage) Get(paths []string) (interface{}, []string) {
	if len(paths) == 0 {
		return p, nil
	}
	x := p.GetPkg(paths[0], false)
	if x != nil {
		return x.Get(paths[1:])
	}
	m := p.GetModel(paths[0])
	if m != nil {
		return m, paths[1:]
	}
	return nil, nil
}

type ModelPackages []*ModelPackage

func ToModelPackage(models Models) *ModelPackage {
	ret := &ModelPackage{Key: "_root"}
	for _, m := range models {
		ret.Add(m.Pkg, m)
	}
	return ret
}

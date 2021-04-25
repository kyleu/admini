package schema

import (
	"github.com/kyleu/admini/app/util"
)

type ModelPackage struct {
	Key           string        `json:"key"`
	Pkg           util.Pkg      `json:"pkg"`
	ChildModels   Models        `json:"childModels,omitempty"`
	ChildPackages ModelPackages `json:"childPackages,omitempty"`
}

func (p *ModelPackage) Path() string {
	return p.Pkg.ToPath(p.Key)
}

func (p *ModelPackage) GetPkg(key string, createIfMissing bool) *ModelPackage {
	for _, x := range p.ChildPackages {
		if x.Key == key {
			return x
		}
	}
	if createIfMissing {
		pkgs := append(p.Pkg, p.Key)
		if pkgs[0] == "_root" {
			pkgs = pkgs[1:]
		}
		x := &ModelPackage{Key: key, Pkg: pkgs}
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
	if x := p.GetPkg(paths[0], false); x != nil {
		return x.Get(paths[1:])
	}
	m := p.GetModel(paths[0])
	if m == nil {
		return p, paths
	}
	return m, paths[1:]
}

type ModelPackages []*ModelPackage

func ToModelPackage(models Models) *ModelPackage {
	ret := &ModelPackage{Key: "_root"}
	for _, m := range models {
		ret.Add(m.Pkg, m)
	}
	return ret
}

package model

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/util"
)

type Relationship struct {
	Key          string   `json:"key"`
	SourceFields []string `json:"srcFields"`
	TargetPkg    util.Pkg `json:"tgtPackage"`
	TargetModel  string   `json:"tgtModel"`
	TargetFields []string `json:"tgtFields"`
}

func (r *Relationship) String() string {
	return fmt.Sprintf("%v: [%v] -> %v[%v]", r.Key, strings.Join(r.SourceFields, ", "), r.TargetModel, strings.Join(r.TargetFields, ", "))
}

type Relationships []*Relationship

func (s Relationships) Get(key string) *Relationship {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func (m *Model) ApplicableRelations(key string) Relationships {
	ret := Relationships{}
	for _, r := range m.Relationships {
		for _, sf := range r.SourceFields {
			if sf == key {
				ret = append(ret, r)
			}
		}
	}
	return ret
}

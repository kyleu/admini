package schema

import (
	"fmt"
	"sort"

	"github.com/kyleu/admini/app/util"
)

type Model struct {
	Key        string    `json:"key"`
	Pkg        util.Pkg  `json:"pkg,omitempty"`
	Type       ModelType `json:"type"`
	Interfaces []string  `json:"interfaces,omitempty"`
	Fields     Fields    `json:"fields,omitempty"`
	Indexes    Indexes   `json:"indexes,omitempty"`
	Metadata   *Metadata `json:"metadata,omitempty"`
}

func (m *Model) String() string {
	if len(m.Pkg) == 0 {
		return m.Key
	}
	return m.Pkg.StringWith(m.Key)
}

func (m *Model) ClassName() string {
	return m.Key
}

func (m *Model) PropName() string {
	return m.Key
}

func (m *Model) AddField(f *Field) error {
	if f == nil {
		return fmt.Errorf("nil field")
	}
	if m.Fields.Get(f.Key) != nil {
		return fmt.Errorf(alreadyExists("field", f.Key))
	}
	m.Fields = append(m.Fields, f)
	m.Fields.Sort()
	return nil
}

func (m *Model) AddIndex(i *Index) error {
	if i == nil {
		return fmt.Errorf("nil index")
	}
	if m.Fields.Get(i.Key) != nil {
		return fmt.Errorf(alreadyExists("index", i.Key))
	}
	m.Indexes = append(m.Indexes, i)
	return nil
}

func (m *Model) Path() string {
	return m.Pkg.ToPath(m.Key)
}

type Models []*Model

func (m Models) Get(pkg util.Pkg, key string) *Model {
	for _, x := range m {
		if x.Pkg.Equals(pkg) && x.Key == key {
			return x
		}
	}
	return nil
}

func (m Models) HasField() bool {
	for _, model := range m {
		if len(model.Fields) > 0 {
			return true
		}
	}
	return false
}

func (m Models) HasIndex() bool {
	for _, model := range m {
		if len(model.Indexes) > 0 {
			return true
		}
	}
	return false
}

func (m Models) ByType(t ModelType) Models {
	ret := Models{}
	for _, x := range m {
		if x.Type == t {
			ret = append(ret, x)
		}
	}
	return ret
}

func (m Models) Sort() {
	sort.Slice(m, func(l int, r int) bool {
		return m[l].Key < m[r].Key
	})
}

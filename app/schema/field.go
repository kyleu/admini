package schema

import (
	"sort"

	"github.com/kyleu/admini/app/schema/schematypes"
)

type Field struct {
	Key      string               `json:"key"`
	Ordinal  int                  `json:"ordinal,omitempty"`
	Type     *schematypes.Wrapped `json:"type"`
	Default  interface{}          `json:"default,omitempty"`
	Nullable bool                 `json:"nullable,omitempty"`
	ReadOnly bool                 `json:"readOnly,omitempty"`
	Metadata *Metadata            `json:"metadata,omitempty"`
}

func (f *Field) String() string {
	return f.Key + " " + f.Type.String()
}

type Fields []*Field

func (s Fields) Get(key string) *Field {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func (s Fields) Sort() {
	sort.Slice(s, func(l int, r int) bool {
		if s[l].Ordinal == s[r].Ordinal {
			return s[l].Key < s[r].Key
		}
		return s[l].Ordinal < s[r].Ordinal
	})
}

package field

import (
	"github.com/kyleu/admini/app/schema/schematypes"
)

type Field struct {
	Key      string               `json:"key"`
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

func (s Fields) Names() []string {
	ret := make([]string, 0, len(s))
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}

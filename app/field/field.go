package field

import (
	"github.com/kyleu/admini/app/schema/schematypes"
)

type Field struct {
	Key      string               `json:"key"`
	Type     *schematypes.Wrapped `json:"type"`
	Default  interface{}          `json:"default,omitempty"`
	ReadOnly bool                 `json:"readOnly,omitempty"`
	Metadata *Metadata            `json:"metadata,omitempty"`
}

func (f *Field) String() string {
	return f.Key + " " + f.Type.String()
}

func (f *Field) Description() string {
	if f.Metadata == nil {
		return ""
	}
	return f.Metadata.Description
}

func (f *Field) Nullable() bool {
	return f.Type.IsOption()
}

type Fields []*Field

func (s Fields) Get(key string) (int, *Field) {
	for idx, x := range s {
		if x.Key == key {
			return idx, x
		}
	}
	return -1, nil
}

func (s Fields) Names() []string {
	ret := make([]string, 0, len(s))
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}

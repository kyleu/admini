package field

import (
	"github.com/kyleu/admini/app/util"
	"time"

	"github.com/kyleu/admini/app/schema/schematypes"
)

type Field struct {
	Key      string               `json:"key"`
	Title    string               `json:"-"`
	Type     *schematypes.Wrapped `json:"type"`
	Default  interface{}          `json:"default,omitempty"`
	ReadOnly bool                 `json:"readOnly,omitempty"`
	Metadata *Metadata            `json:"metadata,omitempty"`
}

func (f *Field) Name() string {
	if f.Title == "" {
		return util.ToTitle(f.Key)
	}
	return f.Title
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

func (f *Field) DefaultClean() interface{} {
	switch f.Default {
	case nil:
		return f.Type.Default(f.Key)
	case "now()":
		return time.Now()
	default:
		return f.Default
	}
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

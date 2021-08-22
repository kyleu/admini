package export

import (
	"fmt"

	"github.com/kyleu/admini/app/util"
)

const (
	LangGo   = "go"
	LangJSON = "json"
)

type Format struct {
	Language string `json:"language"`
	Flavor   string `json:"flavor"`
}

func (f Format) Get(s string) string {
	switch s {
	case "logger":
		switch f.Flavor {
		case util.AppKey:
			return "*zap.Logger"
		default:
			return "*logrus.Logger"
		}
	case "database":
		return "*database.Service"
	default:
		return fmt.Sprintf("unknown key [%s]", s)
	}
}

func (f Format) String() interface{} {
	return fmt.Sprintf("%s:%s", f.Language, f.Flavor)
}

type Formats []*Format

var AllFormats = Formats{
	{Language: LangJSON, Flavor: util.AppKey},
	{Language: LangGo, Flavor: util.AppKey},
	{Language: LangGo, Flavor: "npn"},
	{Language: LangGo, Flavor: "fevo"},
	{Language: LangGo, Flavor: "stock"},
}

func (t Formats) Get(l string, f string) *Format {
	for _, t := range AllFormats {
		if t.Language == l && t.Flavor == f {
			return t
		}
	}
	return nil
}

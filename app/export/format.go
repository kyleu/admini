package export

import (
	"github.com/kyleu/admini/app/util"
)

const LangGo = "go"

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
	default:
		return "unknown key [" + s + "]"
	}
}

type Formats []*Format

var AllFormats = Formats{
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

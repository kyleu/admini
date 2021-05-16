package export

type Format struct {
	Language string `json:"language"`
	Flavor   string `json:"flavor"`
}

const LangGo = "go"

type Formats []*Format

func (t Formats) Get(l string, f string) *Format {
	for _, t := range AllFormats {
		if t.Language == l && t.Flavor == f {
			return t
		}
	}
	return nil
}

var AllFormats = Formats{
	{Language: LangGo, Flavor: "admini"},
	{Language: LangGo, Flavor: "npn"},
	{Language: LangGo, Flavor: "fevo"},
	{Language: LangGo, Flavor: "stock"},
}

type Result struct {
	Key string `json:"key"`
	Out *File  `json:"out"`
}

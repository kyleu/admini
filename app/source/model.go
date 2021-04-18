package source

import "github.com/kyleu/admini/app/schema"

type Source struct {
	Key   string
	Title string
	Paths []string
	Type  schema.Origin
}

type Sources []*Source

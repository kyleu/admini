package source

import (
	"encoding/json"

	"github.com/kyleu/admini/app/schema"
)

type Source struct {
	Key         string          `json:"-"`
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	Paths       []string        `json:"paths,omitempty"`
	Type        schema.Origin   `json:"type,omitempty"`
	Config      json.RawMessage `json:"config,omitempty"`
}

type Sources []*Source

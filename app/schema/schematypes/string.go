package schematypes

import (
	"fmt"
)

const KeyString = "string"

type String struct {
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
}

var _ Type = (*String)(nil)

func (s *String) Key() string {
	return KeyString
}

func (s *String) String() string {
	if s.MaxLength > 0 {
		return fmt.Sprintf("%v(%v)", s.Key(), s.MaxLength)
	}
	return s.Key()
}

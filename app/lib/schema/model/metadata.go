// Package model - Content managed by Project Forge, see [projectforge.md] for details.
package model

type Metadata struct {
	Description string   `json:"description,omitempty"`
	Comments    []string `json:"comments,omitempty"`
	Source      string   `json:"source,omitempty"`
	Line        int      `json:"line,omitempty"`
	Column      int      `json:"column,omitempty"`
}

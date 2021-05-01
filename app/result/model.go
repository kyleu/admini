package result

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/schema"
)

type Timing struct {
	Connected int `json:"connected,omitempty"`
	Prepared  int `json:"prepared,omitempty"`
	Elapsed   int `json:"elapsed,omitempty"`
}

type Result struct {
	Title  string          `json:"title,omitempty"`
	Count  int             `json:"count,omitempty"`
	Query  string          `json:"query,omitempty"`
	Fields schema.Fields   `json:"fields"`
	Data   [][]interface{} `json:"data"`
	Timing *Timing         `json:"timing,omitempty"`
}

func NewResult(title string, count int, q string, fields schema.Fields, data [][]interface{}, timing *Timing) *Result {
	return &Result{
		Title:  title,
		Count:  count,
		Query:  q,
		Fields: fields,
		Data:   data,
		Timing: timing,
	}
}

func (r *Result) Size() int {
	return len(r.Data)
}

func (r *Result) Debug() string {
	sb := &strings.Builder{}
	sb.WriteString(fmt.Sprintf("Result with [%d] fields and [%d] rows", len(r.Fields), len(r.Data)))
	if len(r.Fields) > 0 {
		sb.WriteString(" - ")
	}
	for idx, f := range r.Fields {
		_, _ = sb.WriteString(f.Key + ": " + f.Type.String())
		if idx < len(r.Fields)-1 {
			_, _ = sb.WriteString(", ")
		}
	}
	return sb.String()
}

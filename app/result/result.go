package result

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/schema/field"
)

type Timing struct {
	Connected int `json:"connected,omitempty"`
	Prepared  int `json:"prepared,omitempty"`
	Elapsed   int `json:"elapsed,omitempty"`
}

type Result struct {
	Title    string             `json:"title,omitempty"`
	Count    int                `json:"count,omitempty"`
	Query    string             `json:"query,omitempty"`
	Fields   field.Fields       `json:"fields"`
	Data     [][]any            `json:"data"`
	Timing   *Timing            `json:"timing,omitempty"`
	Children map[string]*Result `json:"children,omitempty"`
}

func NewResult(title string, count int, q string, fields field.Fields, data [][]any, timing *Timing) *Result {
	return &Result{Title: title, Count: count, Query: q, Fields: fields, Data: data, Timing: timing}
}

func (r *Result) Size() int {
	return len(r.Data)
}

func (r *Result) Debug() string {
	sb := &strings.Builder{}
	_, _ = fmt.Fprintf(sb, "Result with [%d] fields and [%d] rows", len(r.Fields), len(r.Data))
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

func (r *Result) SingleRow() (field.Fields, []any, error) {
	if len(r.Data) == 0 {
		return nil, nil, errors.New("empty data")
	}
	if len(r.Data) != 1 {
		return nil, nil, errors.Errorf("required single row, encountered [%d] rows", len(r.Data))
	}
	return r.Fields, r.Data[0], nil
}

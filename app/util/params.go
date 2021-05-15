package util

import (
	"fmt"
	"net/url"
	"strings"

	"go.uber.org/zap"
)

// AllowedColumns A map with arbitrary string keys associated to a string array containing all allowed columns
var AllowedColumns = map[string][]string{}

// Details of a specific set of ordering parameters, with limit and offset
type Params struct {
	Key       string    `json:"key"`
	Orderings Orderings `json:"orderings,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Offset    int       `json:"offset,omitempty"`
}

// Updates or creates Params with the provided Orderings
func ParamsWithDefaultOrdering(key string, params *Params, orderings ...*Ordering) *Params {
	if params == nil {
		params = &Params{Key: key}
	}

	if len(params.Orderings) == 0 {
		params.Orderings = orderings
	}

	return params
}

// Clones this Params, replacing the orderings with the provided arguments
func (p *Params) CloneOrdering(orderings ...*Ordering) *Params {
	if p == nil {
		return nil
	}
	return &Params{Key: p.Key, Orderings: orderings, Limit: p.Limit, Offset: p.Offset}
}

// Indicates if there is more data past the provided page
func (p *Params) HasNextPage(count int) bool {
	if p == nil || p.Limit == 0 {
		return false
	}
	return count > (p.Offset + p.Limit)
}

// Returns a clone of this Params, configured for the next page
func (p *Params) NextPage() *Params {
	limit := p.Limit
	if limit == 0 {
		limit = 100
	}
	offset := p.Offset + limit
	if offset < 0 {
		offset = 0
	}
	return &Params{Key: p.Key, Orderings: p.Orderings, Limit: p.Limit, Offset: offset}
}

// Indicates if there is data prior to the provided page
func (p *Params) HasPreviousPage() bool {
	return p != nil && p.Offset > 0
}

// Returns a clone of this Params, configured for the previous page
func (p *Params) PreviousPage() *Params {
	limit := p.Limit
	if limit == 0 {
		limit = 100
	}
	offset := p.Offset - limit
	if offset < 0 {
		offset = 0
	}
	return &Params{Key: p.Key, Orderings: p.Orderings, Limit: p.Limit, Offset: offset}
}

// Returns the Orderings of this Params that match the provided column
func (p *Params) GetOrdering(col string) *Ordering {
	var ret *Ordering

	for _, o := range p.Orderings {
		if o.Column == col {
			ret = o
		}
	}

	return ret
}

// converts this Params into a SQL order by clause
func (p *Params) OrderByString() string {
	ret := make([]string, 0, len(p.Orderings))

	for _, o := range p.Orderings {
		dir := ""
		if !o.Asc {
			dir = " desc"
		}
		ret = append(ret, o.Column+dir)
	}

	return strings.Join(ret, ", ")
}

// Filters this Params, limiting columns to those matching the AllowedColumns
func (p *Params) Filtered(available []string, logger *zap.SugaredLogger) *Params {
	if available == nil {
		available = AllowedColumns[p.Key]
	}

	if len(available) == 0 {
		logger.Warn("no columns available for [" + p.Key + "]")
	}

	if len(p.Orderings) > 0 {
		allowed := Orderings{}

		for _, o := range p.Orderings {
			containsCol := false
			for _, c := range available {
				if c == o.Column {
					containsCol = true
				}
			}
			if containsCol {
				allowed = append(allowed, o)
			} else {
				const msg = "no column [%v] for [%v] available in allowed columns [%v]"
				logger.Warnf(msg, o.Column, p.Key, strings.Join(available, ", "))
			}
		}

		return &Params{Key: p.Key, Orderings: allowed, Limit: p.Limit, Offset: p.Offset}
	}

	return p
}

// String representation of this Params, mostly used for debugging
func (p *Params) String() string {
	ol := ""
	if p.Offset > 0 {
		ol += fmt.Sprintf("%v/", p.Offset)
	}
	if p.Limit > 0 {
		ol += fmt.Sprint(p.Limit)
	}
	ord := make([]string, 0, len(p.Orderings))
	for _, o := range p.Orderings {
		ord = append(ord, o.String())
	}
	return fmt.Sprintf("%v(%v): %v", p.Key, ol, strings.Join(ord, " / "))
}

// Converts this Params to a URL-safe query string
func (p *Params) ToQueryString(u *url.URL) string {
	if p == nil {
		return ""
	}

	if u == nil {
		return ""
	}

	ret := u.Query()

	delete(ret, p.Key+".o")
	delete(ret, p.Key+".l")
	delete(ret, p.Key+".x")

	for _, o := range p.Orderings {
		s := o.Column

		if !o.Asc {
			s += ".d"
		}

		ret.Add(p.Key+".o", s)
	}

	if p.Limit > 0 {
		ret.Add(p.Key+".l", fmt.Sprint(p.Limit))
	}

	if p.Offset > 0 {
		ret.Add(p.Key+".x", fmt.Sprint(p.Offset))
	}

	return ret.Encode()
}

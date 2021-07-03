package filter

import (
	"strings"

	"github.com/kyleu/admini/app/database"
	"go.uber.org/zap"
)

type Options struct {
	Sort   []string `json:"sort,omitempty"`
	Filter []string `json:"filter,omitempty"`
	Group  []string `json:"group,omitempty"`
	Search string   `json:"search,omitempty"`
	Params *Params  `json:"params,omitempty"`
}

func (o *Options) ToSQL(key string, cols []string, logger *zap.SugaredLogger) string {
	where := ""
	return database.SQLSelect(strings.Join(cols, ", "), key, where, o.Params.OrderByString(), o.Params.Limit, o.Params.Offset)
}

type OptionsMap map[string]*Options

func (m OptionsMap) Get(key string) *Options {
	ret, ok := m[key]
	if ok {
		return ret
	}
	return &Options{Params: &Params{Key: key}}
}

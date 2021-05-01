package lpostgres

import (
	"fmt"
	"github.com/kyleu/admini/app/util"
	"strings"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
)

func (l *Loader) List(source string, cfg []byte, model *schema.Model, params util.ParamSet) (*result.Result, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	q := modelListQuery(model, params.Get(model.Key, model.Fields.Names()))
	rows, err := db.Query(q, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing models for [%v]: %w", model.Key, err)
	}

	timing := &result.Timing{}
	ret, err := NewResult(model.Key, q, timing, rows)
	if err != nil {
		return nil, fmt.Errorf("error constructing result for [%v]: %w", model.Key, err)
	}

	count, err := l.Count(source, cfg, model)
	if err != nil {
		return nil, fmt.Errorf("error constructing result for [%v]: %w", model.Key, err)
	}
	ret.Count = count

	return ret, nil
}

func (l *Loader) Count(source string, cfg []byte, model *schema.Model) (int, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return 0, fmt.Errorf("error opening database: %w", err)
	}

	q := modelCountQuery(model)
	c := struct{
		C int `db:"c"`
	}{}
	err = db.Get(&c, q, nil)
	if err != nil {
		return 0, fmt.Errorf("error listing models for [%v]: %w", model.Key, err)
	}
	return c.C, nil
}

func modelListQuery(m *schema.Model, params *util.Params) string {
	cols := make([]string, 0, len(m.Fields))
	for _, f := range m.Fields {
		def := "\""+f.Key+"\""
		cols = append(cols, def)
	}
	tbl := "\""+m.Key+"\""
	if len(m.Pkg) > 0{
		l := m.Pkg.Last()
		if l != "public" {
			tbl = "\"" + l + "\"." + tbl
		}
	}

	return database.SQLSelect(strings.Join(cols, ", "), tbl, "", params.OrderByString(), params.Limit, params.Offset)
}

func modelCountQuery(m *schema.Model) string {
	tbl := "\""+m.Key+"\""
	if len(m.Pkg) > 0{
		l := m.Pkg.Last()
		if l != "public" {
			tbl = "\"" + l + "\"." + tbl
		}
	}

	return database.SQLSelectSimple("count(*) as c", tbl, "")
}

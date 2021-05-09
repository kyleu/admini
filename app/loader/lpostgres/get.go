package lpostgres

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
)

func (l *Loader) Get(source string, cfg []byte, m *model.Model, ids []interface{}) (*result.Result, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	q, err := modelGetByPKQuery(m)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(q, nil, ids...)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error listing models for [%v]", m.Key))
	}

	var timing *result.Timing
	ret, err := ParseResultFields(m.Key, 0, q, timing, m.Fields, rows)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error constructing result for [%v]", m.Key))
	}

	return ret, nil
}

func modelGetByPKQuery(m *model.Model) (string, error) {
	cols, tbl := forTable(m)
	pk := m.GetPK()
	if len(pk) == 0 {
		return "", errors.New(fmt.Sprintf("no PK for model [%v]", m.Key))
	}
	where := []string{}
	for idx, pkf := range pk {
		where = append(where, fmt.Sprintf(`"%v" = $%v`, pkf, idx+1))
	}
	return database.SQLSelectSimple(cols, tbl, strings.Join(where, " and ")), nil
}

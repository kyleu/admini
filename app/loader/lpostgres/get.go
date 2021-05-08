package lpostgres

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
)

func (l *Loader) Get(source string, cfg []byte, m *model.Model, ids []interface{}) (*result.Result, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	q, err := modelGetByPKQuery(m)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(q, nil, ids...)
	if err != nil {
		return nil, fmt.Errorf("error listing models for [%v]: %w", m.Key, err)
	}

	var timing *result.Timing
	ret, err := ParseResultFields(m.Key, 0, q, timing, m.Fields, rows)
	if err != nil {
		return nil, fmt.Errorf("error constructing result for [%v]: %w", m.Key, err)
	}

	return ret, nil
}

func modelGetByPKQuery(m *model.Model) (string, error) {
	cols, tbl := forTable(m)
	pk := m.GetPK()
	if len(pk) == 0 {
		return "", fmt.Errorf("no PK for model [%v]", m.Key)
	}
	where := []string{}
	for idx, pkf := range pk {
		where = append(where, fmt.Sprintf(`"%v" = $%v`, pkf, idx+1))
	}
	return database.SQLSelectSimple(cols, tbl, strings.Join(where, " and ")), nil
}

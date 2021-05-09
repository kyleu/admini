package postgres

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries"
)

type tableResult struct {
	Schema string `db:"schema"`
	Name   string `db:"name"`
	Type   string `db:"type"`
	Owner  string `db:"owner"`
}

func (t tableResult) ToModel() *model.Model {
	ret := &model.Model{
		Key: t.Name,
		Pkg: util.Pkg{t.Schema},
	}

	switch t.Type {
	case "table":
		ret.Type = model.ModelTypeStruct
	case "view":
		ret.Type = model.ModelTypeInterface
	case "sequence":
		ret.Type = model.ModelTypeSequence
	default:
		util.LogWarn("unknown model type [" + t.Type + "]")
		ret.Type = model.ModelTypeUnknown
	}
	return ret
}

func loadTables(db *database.Service) (model.Models, error) {
	tables := []*tableResult{}
	err := db.Select(&tables, queries.ListTables(db.SchemaName), nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't list tables")
	}

	util.LogInfo(fmt.Sprintf("loading [%v] tables", len(tables)))

	ret := make(model.Models, 0, len(tables))
	for _, t := range tables {
		ret = append(ret, t.ToModel())
	}
	ret.Sort()
	return ret, nil
}

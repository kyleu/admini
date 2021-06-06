package postgres

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries/qpostgres"
)

type indexResult struct {
	Schema      string `db:"schema_name"`
	Table       string `db:"table_name"`
	Index       string `db:"index_name"`
	PrimaryKey  bool   `db:"pk"`
	Unique      bool   `db:"u"`
	ColumnNames string `db:"column_names"`
}

func (r indexResult) AsIndex() *model.Index {
	return &model.Index{
		Key:      r.Index,
		Fields:   util.SplitAndTrim(r.ColumnNames, ","),
		Unique:   r.Unique,
		Primary:  r.PrimaryKey,
		Metadata: nil,
	}
}

func loadIndexes(models model.Models, db *database.Service) error {
	var idxs []*indexResult
	err := db.Select(&idxs, qpostgres.ListIndexes(db.SchemaName), nil)
	if err != nil {
		return errors.Wrap(err, "can't list indexes")
	}

	for _, idx := range idxs {
		mod := models.Get(util.Pkg{idx.Schema}, idx.Table)
		if mod == nil {
			return errors.Errorf("no table [%s] found among [%d] candidates", idx.Table, len(models))
		}
		err = mod.AddIndex(idx.AsIndex())
		if err != nil {
			return errors.Wrap(err, "can't add field")
		}
	}
	return nil
}

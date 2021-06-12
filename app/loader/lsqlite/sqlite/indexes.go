package sqlite

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/queries/qsqlite"
)

type indexResult struct {
	Table       string `db:"xn"`
	Name        string `db:"n"`
	ColumnIndex string `db:"ci"`
	ColumnName  string `db:"cn"`
}

func loadIndexes(models model.Models, db *database.Service) error {
	var idxs []*indexResult
	err := db.Select(&idxs, qsqlite.ListIndexes(db.SchemaName), nil)
	if err != nil {
		return errors.Wrap(err, "can't list indexes")
	}

	for _, idx := range idxs {
		mod := models.Get(nil, idx.Table)
		if mod == nil {
			return errors.Errorf("no table [%s] found among [%d] candidates", idx.Table, len(models))
		}

		curr := mod.Indexes.Get(idx.Name)
		if curr == nil {
			curr = &model.Index{Key: idx.Name}
			err = mod.AddIndex(curr)
			if err != nil {
				return errors.Wrap(err, "can't add index")
			}
		}

		curr.Fields = append(curr.Fields, idx.ColumnName)
	}
	return nil
}

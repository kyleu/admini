package sqlite

import (
	"context"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qsqlite"
)

type indexResult struct {
	Table       string `db:"xn"`
	Name        string `db:"n"`
	ColumnIndex string `db:"ci"`
	ColumnName  string `db:"cn"`
}

func loadIndexes(ctx context.Context, models model.Models, db *database.Service, logger util.Logger) error {
	var idxs []*indexResult
	err := db.Select(ctx, &idxs, qsqlite.ListIndexes(db.SchemaName), nil, logger)
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

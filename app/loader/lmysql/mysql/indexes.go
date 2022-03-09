package mysql

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qmysql"
)

type indexResult struct {
	Schema      string `db:"schema_name"`
	Table       string `db:"table_name"`
	Index       string `db:"index_name"`
	NonUnique   bool   `db:"non_unique"`
	ColumnNames string `db:"column_names"`
}

func (r indexResult) AsIndex() *model.Index {
	return &model.Index{
		Key:      r.Index,
		Fields:   util.StringSplitAndTrim(r.ColumnNames, ","),
		Unique:   !r.NonUnique,
		Primary:  strings.EqualFold(r.Index, "primary"),
		Metadata: nil,
	}
}

func loadIndexes(ctx context.Context, models model.Models, db *database.Service) error {
	var idxs []*indexResult
	err := db.Select(ctx, &idxs, qmysql.ListIndexes(db.DatabaseName), nil)
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
			return errors.Wrap(err, "can't add index")
		}
	}
	return nil
}

package sqlserver

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qsqlserver"
)

type indexResult struct {
	Schema      string         `db:"schema_name"`
	Table       string         `db:"table_name"`
	Index       sql.NullString `db:"index_name"`
	PrimaryKey  bool           `db:"pk"`
	Unique      bool           `db:"u"`
	ColumnNames sql.NullString `db:"column_names"`
}

func (r indexResult) AsIndex() *model.Index {
	return &model.Index{
		Key:      r.Index.String,
		Fields:   util.StringSplitAndTrim(r.ColumnNames.String, ","),
		Unique:   r.Unique,
		Primary:  r.PrimaryKey,
		Metadata: nil,
	}
}

func loadIndexes(ctx context.Context, models model.Models, db *database.Service, logger util.Logger) error {
	var idxs []*indexResult
	err := db.Select(ctx, &idxs, qsqlserver.ListIndexes(db.SchemaName), nil, logger)
	if err != nil {
		return errors.Wrap(err, "can't list indexes")
	}

	for _, idx := range idxs {
		if !idx.Index.Valid || !idx.ColumnNames.Valid {
			continue
		}
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

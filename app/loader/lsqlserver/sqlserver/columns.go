package sqlserver

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qsqlserver"
)

type columnResult struct {
	Schema     string         `db:"schema_name"`
	Table      string         `db:"table_name"`
	Name       string         `db:"column_name"`
	Ordinal    int32          `db:"column_id"`
	Default    sql.NullString `db:"column_default"`
	Nullable   string         `db:"is_nullable"`
	DataType   string         `db:"data_type"`
	CharLength sql.NullInt32  `db:"max_length"`
	NumLength  sql.NullInt32  `db:"precision"`
}

func (cr *columnResult) IsNullable() bool {
	return cr.Nullable == sqlServerYes
}

func (cr *columnResult) AsField(readOnlyOverride bool, enums model.Models, logger util.Logger) *field.Field {
	var d any
	if cr.Default.Valid {
		d = cr.Default.String
	}
	return &field.Field{
		Key:      cr.Name,
		Type:     typeFor(cr.DataType, cr, enums, logger),
		Default:  d,
		ReadOnly: readOnlyOverride,
		Metadata: nil,
	}
}

func loadColumns(ctx context.Context, models model.Models, db *database.Service, logger util.Logger) error {
	var cols []*columnResult
	err := db.Select(ctx, &cols, qsqlserver.ListColumns(db.SchemaName), nil, logger)
	if err != nil {
		return errors.Wrap(err, "can't list columns")
	}

	for _, col := range cols {
		mod := models.Get(util.Pkg{col.Schema}, col.Table)
		if mod == nil {
			return errors.Errorf("no table [%s] found among [%d] candidates", col.Table, len(models))
		}
		err = mod.AddField(col.AsField(mod.Type == model.TypeInterface, models, logger))
		if err != nil {
			return errors.Wrap(err, "can't add field")
		}
	}

	return nil
}

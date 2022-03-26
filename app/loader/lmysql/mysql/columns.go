package mysql

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qmysql"
)

type columnResult struct {
	Schema            string         `db:"TABLE_SCHEMA"`
	Table             string         `db:"TABLE_NAME"`
	Name              string         `db:"COLUMN_NAME"`
	Ordinal           int32          `db:"ORDINAL_POSITION"`
	Default           sql.NullString `db:"COLUMN_DEFAULT"`
	Nullable          string         `db:"IS_NULLABLE"`
	DataType          string         `db:"DATA_TYPE"`
	CharLength        sql.NullInt64  `db:"CHARACTER_MAXIMUM_LENGTH"`
	OctetLength       sql.NullInt64  `db:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision  sql.NullInt64  `db:"NUMERIC_PRECISION"`
	NumericScale      sql.NullInt64  `db:"NUMERIC_SCALE"`
	DatetimePrecision sql.NullInt64  `db:"DATETIME_PRECISION"`
}

func (cr *columnResult) AsField(readOnlyOverride bool, logger *zap.SugaredLogger) *field.Field {
	var d any
	if cr.Default.Valid {
		d = cr.Default.String
	}
	return &field.Field{
		Key:      cr.Name,
		Type:     typeFor(cr.DataType, cr, logger),
		Default:  d,
		ReadOnly: readOnlyOverride,
		Metadata: nil,
	}
}

func loadColumns(ctx context.Context, models model.Models, db *database.Service, logger *zap.SugaredLogger) error {
	var cols []*columnResult
	err := db.Select(ctx, &cols, qmysql.ListColumns(db.DatabaseName), nil, logger)
	if err != nil {
		return errors.Wrap(err, "can't list columns")
	}

	for _, col := range cols {
		mod := models.Get(util.Pkg{col.Schema}, col.Table)
		if mod == nil {
			return errors.Errorf("no table [%s] found among [%d] candidates (%s)", col.Table, len(models), util.StringArrayOxfordComma(models.Names(), "and"))
		}
		err = mod.AddField(col.AsField(mod.Type == model.TypeInterface, logger))
		if err != nil {
			return errors.Wrap(err, "can't add field")
		}
	}

	return nil
}

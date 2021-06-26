package postgres

import (
	"database/sql"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/field"
	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries/qpostgres"
)

type columnResult struct {
	Schema                string         `db:"table_schema"`
	Table                 string         `db:"table_name"`
	Name                  string         `db:"column_name"`
	Ordinal               int32          `db:"ordinal_position"`
	Default               sql.NullString `db:"column_default"`
	Nullable              string         `db:"is_nullable"`
	DataType              string         `db:"data_type"`
	ArrayType             sql.NullString `db:"array_type"`
	CharLength            sql.NullInt32  `db:"character_maximum_length"`
	OctetLength           sql.NullInt32  `db:"character_octet_length"`
	NumericPrecision      sql.NullInt32  `db:"numeric_precision"`
	NumericPrecisionRadix sql.NullInt32  `db:"numeric_precision_radix"`
	NumericScale          sql.NullInt32  `db:"numeric_scale"`
	DatetimePrecision     sql.NullInt32  `db:"datetime_precision"`
	IntervalType          sql.NullInt32  `db:"interval_type"`
	DomainSchema          sql.NullString `db:"domain_schema"`
	DomainName            sql.NullString `db:"domain_name"`
	UDTSchema             string         `db:"udt_schema"`
	UDTName               string         `db:"udt_name"`
	DTDIdentifier         string         `db:"dtd_identifier"`
	Updatable             string         `db:"is_updatable"`
}

func (cr *columnResult) IsNullable() bool {
	return cr.Nullable == pgYes
}

func (cr *columnResult) AsField(readOnlyOverride bool, logger *zap.SugaredLogger) *field.Field {
	var d interface{}
	if cr.Default.Valid {
		d = cr.Default.String
	}
	return &field.Field{
		Key:      cr.Name,
		Type:     typeFor(cr.UDTName, cr, logger),
		Default:  d,
		ReadOnly: readOnlyOverride || (cr.Updatable == pgNo),
		Metadata: nil,
	}
}

func loadColumns(models model.Models, db *database.Service, logger *zap.SugaredLogger) error {
	var cols []*columnResult
	err := db.Select(&cols, qpostgres.ListColumns(db.SchemaName), nil)
	if err != nil {
		return errors.Wrap(err, "can't list columns")
	}

	for _, col := range cols {
		mod := models.Get(util.Pkg{col.Schema}, col.Table)
		if mod == nil {
			return errors.Errorf("no table [%s] found among [%d] candidates", col.Table, len(models))
		}
		err = mod.AddField(col.AsField(mod.Type == model.TypeInterface, logger))
		if err != nil {
			return errors.Wrap(err, "can't add field")
		}
	}

	return nil
}
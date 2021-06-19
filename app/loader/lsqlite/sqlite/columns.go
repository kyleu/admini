package sqlite

import (
	"database/sql"
	"sort"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/field"
	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/queries/qsqlite"
)

type columnResult struct {
	Table    string         `db:"xn"`
	Ordinal  int32          `db:"i"`
	Name     string         `db:"n"`
	DataType string         `db:"t"`
	PK       int            `db:"pk"`
	Default  sql.NullString `db:"dv"`
	NotNull  int            `db:"nn"`
}

func (cr *columnResult) IsNullable() bool {
	return cr.NotNull == 0
}

func (cr *columnResult) AsField(readOnlyOverride bool, logger *zap.SugaredLogger) *field.Field {
	var d interface{}
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

func loadColumns(models model.Models, db *database.Service, logger *zap.SugaredLogger) error {
	var cols []*columnResult
	err := db.Select(&cols, qsqlite.ListColumns(db.SchemaName), nil)
	if err != nil {
		return errors.Wrap(err, "can't list columns")
	}

	pks := map[*model.Model]map[int]string{}

	for _, col := range cols {
		mod := models.Get(nil, col.Table)
		if mod == nil {
			return errors.Errorf("no table [%s] found among [%d] candidates", col.Table, len(models))
		}
		err = mod.AddField(col.AsField(mod.Type == model.TypeInterface, logger))
		if err != nil {
			return errors.Wrap(err, "can't add field")
		}
		if col.PK > 0 {
			curr, ok := pks[mod]
			if !ok {
				curr = map[int]string{}
				pks[mod] = curr
			}
			curr[col.PK] = col.Name
		}
	}

	for k, v := range pks {
		idx := &model.Index{Key: k.Key + "_pk", Unique: true, Primary: true}
		nums := make([]int, 0, len(v))
		for num := range v {
			nums = append(nums, num)
		}
		sort.Ints(nums)
		for _, num := range nums {
			idx.Fields = append(idx.Fields, v[num])
		}
		k.Indexes = append(k.Indexes, idx)
	}

	return nil
}

package mysql

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qmysql"
)

type foreignKeyResult struct {
	Name         string `db:"constraint_name"`
	Ordinal      int    `db:"ordinal"`
	Schema       string `db:"schema_name"`
	Table        string `db:"table_name"`
	Field        string `db:"column_name"`
	TargetSchema string `db:"foreign_schema_name"`
	TargetTable  string `db:"foreign_table_name"`
	TargetField  string `db:"foreign_column_name"`
}

type foreignKeyResults []*foreignKeyResult

func (r foreignKeyResults) Sort() {
	sort.Slice(r, func(i int, j int) bool {
		if r[i].Name == r[j].Name {
			return r[i].Ordinal < r[j].Ordinal
		}
		return r[i].Name < r[j].Name
	})
}

func loadForeignKeys(ctx context.Context, models model.Models, db *database.Service, logger util.Logger) error {
	keys := foreignKeyResults{}
	err := db.Select(ctx, &keys, qmysql.ListForeignKeys(db.DatabaseName), nil, logger)
	if err != nil {
		return errors.Wrap(err, "can't list foreign keys")
	}

	for _, k := range keys {
		mod := models.Get(util.Pkg{k.Schema}, k.Table)
		if mod == nil {
			return errors.Errorf("no model [%s] found among [%d] candidates", k.Table, len(models))
		}

		curr := mod.Relationships.Get(k.Name)
		if curr == nil {
			curr = &model.Relationship{
				Key:          k.Name,
				SourceFields: []string{},
				TargetPkg:    util.Pkg{k.TargetSchema},
				TargetModel:  k.TargetTable,
				TargetFields: []string{},
			}
			mod.Relationships = append(mod.Relationships, curr)
		}

		curr.SourceFields = append(curr.SourceFields, k.Field)
		curr.TargetFields = append(curr.TargetFields, k.TargetField)
	}

	return nil
}

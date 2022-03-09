package postgres

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qpostgres"
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

func loadForeignKeys(ctx context.Context, models model.Models, db *database.Service) error {
	keys := foreignKeyResults{}
	err := db.Select(ctx, &keys, qpostgres.ListForeignKeys(db.SchemaName), nil)
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

func loadEnumRelations(ctx context.Context, enums model.Models, models model.Models, db *database.Service) error {
	for _, e := range enums {
		for _, m := range models {
			for _, f := range m.Fields {
				if ek := f.Type.EnumKey(); ek == e.Key {
					rel := &model.Relationship{
						Key:          "enum_" + m.Key + "_" + f.Key,
						SourceFields: []string{e.Key},
						TargetPkg:    m.Pkg,
						TargetModel:  m.Key,
						TargetFields: []string{f.Key},
					}
					e.Relationships = append(e.Relationships, rel)
				}
			}
		}
	}

	return nil
}

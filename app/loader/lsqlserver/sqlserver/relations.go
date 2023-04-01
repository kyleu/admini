package sqlserver

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qsqlserver"
)

type foreignKeyResult struct {
	Name        string `db:"fk"`
	Schema      string `db:"sch"`
	Table       string `db:"tbl"`
	Field       string `db:"col"`
	TargetTable string `db:"ref_tbl"`
	TargetField string `db:"ref_col"`
}

type foreignKeyResults []*foreignKeyResult

func (r foreignKeyResults) Sort() {
	sort.Slice(r, func(i int, j int) bool {
		return r[i].Name < r[j].Name
	})
}

func loadForeignKeys(ctx context.Context, models model.Models, db *database.Service, logger util.Logger) error {
	keys := foreignKeyResults{}
	err := db.Select(ctx, &keys, qsqlserver.ListForeignKeys(db.SchemaName), nil, logger)
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
				TargetPkg:    util.Pkg{k.Schema},
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

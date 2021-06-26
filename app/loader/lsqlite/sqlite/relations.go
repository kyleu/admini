package sqlite

import (
	"fmt"
	"sort"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/queries/qsqlite"
)

type foreignKeyResult struct {
	Ordinal     int    `db:"idx"`
	Table       string `db:"src"`
	Field       string `db:"src_col"`
	TargetTable string `db:"tgt"`
	TargetField string `db:"tgt_col"`
}

func (r *foreignKeyResult) String() string {
	return fmt.Sprintf("rel-%s-%s", r.Table, r.TargetTable)
}

type foreignKeyResults []*foreignKeyResult

func (r foreignKeyResults) Sort() {
	sort.Slice(r, func(i int, j int) bool {
		if r[i].Table == r[j].Table {
			return r[i].Ordinal < r[j].Ordinal
		}
		return r[i].Table < r[j].Table
	})
}

func loadForeignKeys(models model.Models, db *database.Service) error {
	keys := foreignKeyResults{}
	err := db.Select(&keys, qsqlite.ListForeignKeys(db.SchemaName), nil)
	if err != nil {
		return errors.Wrap(err, "can't list foreign keys")
	}

	for _, k := range keys {
		mod := models.Get(nil, k.Table)
		if mod == nil {
			return errors.Errorf("no model [%s] found among [%d] candidates", k.Table, len(models))
		}

		curr := mod.Relationships.Get(k.String())
		if curr == nil {
			curr = &model.Relationship{
				Key:          k.String(),
				SourceFields: []string{},
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
package postgres

import (
	"fmt"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries"
)

type indexResult struct {
	Schema      string `db:"schema_name"`
	Table       string `db:"table_name"`
	Index       string `db:"index_name"`
	PrimaryKey  bool   `db:"pk"`
	Unique      bool   `db:"u"`
	ColumnNames string `db:"column_names"`
}

func (r indexResult) AsIndex() *schema.Index {
	return &schema.Index{
		Key:      r.Index,
		Fields:   util.SplitAndTrim(r.ColumnNames, ","),
		Unique:   r.Unique,
		Primary:  r.PrimaryKey,
		Metadata: nil,
	}
}

func loadIndexes(models schema.Models, db *database.Service) error {
	idxs := []*indexResult{}
	err := db.Select(&idxs, queries.ListIndexes(), nil)
	if err != nil {
		return fmt.Errorf("can't list indexes: %w", err)
	}

	for _, idx := range idxs {
		mod := models.Get(util.Pkg{idx.Schema}, idx.Table)
		if mod == nil {
			return fmt.Errorf("no table [%v] found among [%v] candidates", idx.Table, len(models))
		}
		err = mod.AddIndex(idx.AsIndex())
		if err != nil {
			return fmt.Errorf("can't add field: %w", err)
		}
	}
	return nil
}

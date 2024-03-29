package sqlserver

import (
	"context"
	"database/sql"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/field"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/lib/types"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qsqlserver"
)

type enumResult struct {
	Schema      string         `db:"schema"`
	Name        string         `db:"name"`
	Internal    string         `db:"internal"`
	Size        string         `db:"size"`
	Elements    string         `db:"elements"`
	Owner       sql.NullString `db:"owner"`
	Privileges  sql.NullString `db:"privileges"`
	Description sql.NullString `db:"description"`
}

func (t *enumResult) ToModel() *model.Model {
	ret := model.NewModel(util.Pkg{t.Schema}, t.Name)
	ret.Type = model.TypeEnum

	els := strings.Split(t.Elements, "\n")
	fields := make(field.Fields, 0, len(els))
	for _, el := range els {
		fields = append(fields, &field.Field{
			Key:  el,
			Type: types.NewEnumValue(),
		})
	}
	ret.Fields = fields

	return ret
}

func loadEnums(ctx context.Context, db *database.Service, logger util.Logger) (model.Models, error) {
	var enums []*enumResult
	err := db.Select(ctx, &enums, qsqlserver.ListTypes(db.SchemaName), nil, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't list enums")
	}

	logger.Infof("loading [%d] enums", len(enums))

	ret := make(model.Models, 0, len(enums))
	for _, t := range enums {
		ret = append(ret, t.ToModel())
	}
	return ret, nil
}

package postgres

import (
	"database/sql"
	"strings"

	"github.com/kyleu/admini/app/field"
	"github.com/kyleu/admini/app/schema/schematypes"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries/qpostgres"
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
			Type: schematypes.NewEnumValue(),
		})
	}
	ret.Fields = fields

	return ret
}

func loadEnums(db *database.Service, logger *zap.SugaredLogger) (model.Models, error) {
	var enums []*enumResult
	err := db.Select(&enums, qpostgres.ListTypes(db.SchemaName), nil)
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

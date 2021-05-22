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
	"github.com/kyleu/admini/queries"
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

func (t enumResult) ToModel(logger *zap.SugaredLogger) *model.Model {
	els := strings.Split(t.Elements, "\n")
	fields := make(field.Fields, 0, len(els))
	for _, el := range els {
		fields = append(fields, &field.Field{
			Key:  el,
			Type: schematypes.NewEnumValue(),
		})
	}

	ret := &model.Model{
		Key:    t.Name,
		Type:   model.ModelTypeEnum,
		Fields: fields,
		Pkg:    util.Pkg{t.Schema},
	}
	return ret
}

func loadEnums(db *database.Service, logger *zap.SugaredLogger) (model.Models, error) {
	enums := []*enumResult{}
	err := db.Select(&enums, queries.ListTypes(db.SchemaName), nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't list enums")
	}

	logger.Infof("loading [%v] enums", len(enums))

	ret := make(model.Models, 0, len(enums))
	for _, t := range enums {
		ret = append(ret, t.ToModel(logger))
	}
	return ret, nil
}

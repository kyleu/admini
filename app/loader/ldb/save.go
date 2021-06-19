package ldb

import (
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Save(db *database.Service, m *model.Model, changes util.ValueMap, logger *zap.SugaredLogger) ([]interface{}, error) {
	return nil, errors.New("TODO: add not implemented")
}



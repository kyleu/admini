package export

import (
	"github.com/kyleu/admini/app/schema/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Result struct {
	Key string `json:"key"`
	Out *File  `json:"out"`
}

type Results []*Result

func Export(m *model.Model, logger *zap.SugaredLogger) (Results, error) {
	svcFile, err := fileService(m, logger)
	if err != nil {
		return nil, errors.Wrapf(err, "error processing service file for [%s]", m.Key)
	}
	modelFile, err := fileModel(m, logger)
	if err != nil {
		return nil, errors.Wrapf(err, "error processing model file for [%s]", m.Key)
	}
	return Results{modelFile, svcFile}, nil
}

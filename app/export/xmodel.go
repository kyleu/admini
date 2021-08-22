package export

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"
	"go.uber.org/zap"
)

func Model(m *model.Model, t *Format, logger *zap.SugaredLogger) (Results, error) {
	switch t.Language {
	case "go":
		svcFile, err := goServiceFile(m, t)
		if err != nil {
			return nil, errors.Wrapf(err, "error processing [%s] service file for [%s]", t.String(), m.Key)
		}
		modelFile, err := goModelFile(m, t, logger)
		if err != nil {
			return nil, errors.Wrapf(err, "error processing [%s] model file for [%s]", t.String(), m.Key)
		}
		return Results{modelFile, svcFile}, nil
	case "json":
		return Results{jsonFile(m, t, logger)}, nil
	default:
		return nil, errors.Errorf("unhandled language [%s]", t.Language)
	}
}

package export

import (
	"github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

func jsonFile(m *model.Model, fm *Format, logger *zap.SugaredLogger) *Result {
	f := NewJSONFile(m.Pkg, m.Key)
	json := util.ToJSON(m)
	f.W(json)
	return &Result{Key: "json", Out: f}
}

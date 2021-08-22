package export

import (
	"strings"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func goServiceFile(m *model.Model, fm *Format) (*Result, error) {
	f := NewGoFile(m.Pkg, m.Key+"Service")

	f.W("type Service struct {", 1)
	f.W("logger " + fm.Get("logger"))
	f.W("}", -1)
	f.W("")
	f.W("func NewService(logger "+fm.Get("logger")+") *Service {", 1)
	f.W("return &Service{logger: logger}")
	f.W("}", -1)
	f.W("")
	f.W("func (s *Service) Test() bool {", 1)
	f.W("return true")
	f.W("}", -1)

	pk := m.GetPK(nil)
	if len(pk) > 0 {
		f.W("")
		pkArgs := make([]string, 0, len(pk))
		for _, pkCol := range pk {
			_, fld := m.Fields.Get(pkCol)
			if fld == nil {
				return nil, errors.Errorf("missing pk col [%s]", pkCol)
			}
			flk := util.ToLowerCamel(util.ToSingular(fld.Key))
			typ, imps := typeString(fld.Type, fm, "model")
			for _, imp := range imps {
				f.AddImport(imp.String())
			}
			pkArgs = append(pkArgs, flk+ " " + typ)
		}
		sk := util.ToCamel(util.ToSingular(m.Key))
		f.AddImport("context")
		f.W("func (s *Service) Get(ctx context.Context, " + strings.Join(pkArgs, ", ") + ") (*" + sk + ", error) {", 1)
		f.W("return true")
		f.W("}", -1)
	}

	ret := &Result{Key: "service", Out: f}
	return ret, nil
}

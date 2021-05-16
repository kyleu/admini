package export

import (
	"fmt"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

func Model(m *model.Model, t *Format, logger *zap.SugaredLogger) ([]*Result, error) {
	return []*Result{goModelFile(m, t, logger), goServiceFile(m, t)}, nil
}

func goModelFile(m *model.Model, fm *Format, logger *zap.SugaredLogger) *Result {
	f := NewGoFile(m.Pkg, m.Key)

	pk := m.GetPK(logger)
	maxKeyLength := 0
	maxTypeLength := 0
	for _, fld := range m.Fields {
		if len(fld.Key) > maxKeyLength {
			maxKeyLength = len(fld.Key)
		}
		x, _ := typeString(fld.Type, fm)
		if len(x) > maxTypeLength {
			maxTypeLength = len(x)
		}
	}

	f.W("type "+util.ToCamel(m.Key)+" struct {", 1)
	msg := "%-" + fmt.Sprintf("%v", maxKeyLength) + "v %-" + fmt.Sprintf("%v", maxTypeLength) + "v %v%v"
	for _, fld := range m.Fields {
		typ, imp := typeString(fld.Type, fm)
		if len(imp) > 0 {
			logger.Warn("TODO: imports!")
		}
		omit := ""
		if fld.Nullable() {
			omit = ",omitempty"
		}
		suffix := ""
		if util.StringArrayContains(pk, fld.Key) {
			suffix = " /* primary key */"
		}
		f.W(fmt.Sprintf(msg, util.ToCamel(fld.Key), typ, "`json:\""+util.ToLowerCamel(fld.Key)+omit+"\"`", suffix))
	}
	f.W("}", -1)

	ret := &Result{Key: "model", Out: f}
	return ret
}

func goServiceFile(m *model.Model, t *Format) *Result {
	f := NewGoFile(m.Pkg, m.Key+"Service")

	f.W("type Service struct {}")
	f.W("")
	f.W("func NewService() *Service {", 1)
	f.W("return &Service()")
	f.W("}", -1)
	f.W("")
	f.W("func (s *Service) Test() bool {", 1)
	f.W("return true")
	f.W("}", -1)

	ret := &Result{Key: "service", Out: f}
	return ret
}

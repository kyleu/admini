package export

import (
	"github.com/kyleu/admini/app/model"
)

func goServiceFile(m *model.Model, t *Format) *Result {
	f := NewGoFile(m.Pkg, m.Key+"Service")

	f.W("type Service struct {", 1)
	f.W("logger " + t.Get("logger"))
	f.W("}", -1)
	f.W("")
	f.W("func NewService(logger "+t.Get("logger")+") *Service {", 1)
	f.W("return &Service{logger: logger}")
	f.W("}", -1)
	f.W("")
	f.W("func (s *Service) Test() bool {", 1)
	f.W("return true")
	f.W("}", -1)

	ret := &Result{Key: "service", Out: f}
	return ret
}

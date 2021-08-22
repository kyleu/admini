package export

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func goServiceFile(m *model.Model, fm *Format) (*Result, error) {
	f := NewGoFile(m.Pkg, m.Key+"Service")

	f.W("type Service struct {", 1)
	f.W("db     " + fm.Get("database"))
	f.W("logger " + fm.Get("logger"))
	f.W("}", -1)
	f.LB()
	nsLine := "func NewService(db %s, logger %s) *Service {"
	f.W(fmt.Sprintf(nsLine, fm.Get("database"), fm.Get("logger")), 1)
	f.W("return &Service{db: db, logger: logger}")
	f.W("}", -1)
	f.LB()
	f.W("func (s *Service) Test() bool {", 1)
	f.W("return true")
	f.W("}", -1)

	pk := m.GetPK(nil)
	if len(pk) > 0 {
		f.LB()
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
		lk := util.ToLowerCamel(util.ToSingular(m.Key))
		sk := util.ToCamel(util.ToSingular(m.Key))
		f.AddImport("context")
		f.W("func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, " + strings.Join(pkArgs, ", ") + ") (*" + sk + ", error) {", 1)
		f.Wf("ret := &%sDTO{}", lk)
		var pkWhereClause []string
		for idx, x := range pk {
			pkWhereClause = append(pkWhereClause, fmt.Sprintf("%s = $%d", x, idx + 1))
		}
		f.Wf("q := database.SQLSelectSimple(Table, ColumnsString, \"%s\")", strings.Join(pkWhereClause, " and "))
		f.Wf("err := s.db.Get(ctx, ret, q, tx, %s)", strings.Join(pk, ", "))
		f.W("if err != nil {", 1)
		var pkLogs []string
		for _, x := range pk {
			pkLogs = append(pkLogs, x + " [%v]")
		}
		msg := "\"unable to get " + lk + " by " + strings.Join(pkLogs, " and ") + "\""
		f.Wf("return nil, errors.Wrapf(err, %s, %s)", msg, strings.Join(pk, ", "))
		f.W("}", -1)
		f.Wf("return ret.To%s(), nil", sk)
		f.W("}", -1)
	}

	ret := &Result{Key: "service", Out: f}
	return ret, nil
}

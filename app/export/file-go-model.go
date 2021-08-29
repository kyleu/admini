package export

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

func goModelFile(m *model.Model, fm *Format, logger *zap.SugaredLogger) (*Result, error) {
	f := NewGoFile(m.Pkg, m.Key)

	pk := m.GetPK(logger)

	lk := util.ToLowerCamel(util.ToSingular(m.Key))
	sk := util.ToCamel(util.ToSingular(m.Key))
	pluralk := util.ToCamel(util.ToPlural(m.Key))
	firstChar := strings.ToLower(string(m.Key[0]))

	maxKeyLength := 0
	maxTypeLength := 0
	maxDTOTypeLength := 0
	dataFields := make([]string, 0, len(m.Fields))
	for _, fld := range m.Fields {
		if len(fld.Key) > maxKeyLength {
			maxKeyLength = len(fld.Key)
		}
		x, _ := typeString(fld.Type, fm, "model")
		if len(x) > maxTypeLength {
			maxTypeLength = len(x)
		}
		z, _ := typeString(fld.Type, fm, "dto")
		if len(z) > maxDTOTypeLength {
			maxDTOTypeLength = len(z)
		}
		dataFields = append(dataFields, firstChar+"."+util.ToCamel(fld.Key))
	}

	f.W("type "+sk+" struct {", 1)
	msg := "%-" + fmt.Sprintf("%d", maxKeyLength) + "s %-" + fmt.Sprintf("%d", maxTypeLength) + "s %s%s"
	for _, fld := range m.Fields {
		typ, imps := typeString(fld.Type, fm, "model")
		for _, imp := range imps {
			f.AddImport(imp.String())
		}
		omit := ""
		if fld.Type.IsOption() {
			omit = ",omitempty"
		}
		suffix := ""
		if util.StringArrayContains(pk, fld.Key) {
			suffix = " /* primary key */"
		}
		f.Wf(msg, util.ToCamel(fld.Key), typ, "`json:\""+util.ToLowerCamel(fld.Key)+omit+"\"`", suffix)
	}
	f.W("}", -1)
	f.LB()

	f.W("func ("+firstChar+" *"+sk+") ToData() []interface{} {", 1)
	f.W("return []interface{}{" + strings.Join(dataFields, ", ") + "}")
	f.W("}", -1)
	f.LB()

	f.Wf("type %s []*%s", pluralk, sk)
	f.LB()

	cols := make([]string, 0, len(m.Fields))
	for _, f := range m.Fields {
		cols = append(cols, "\""+f.Key+"\"")
	}

	f.W("var (", 1)
	f.Wf("Table = \"%s\"", m.Key)
	f.Wf("Columns = []string{%s}", strings.Join(cols, ", "))
	f.AddImport("strings")
	f.Wf("ColumnsString = strings.Join(%sColumns, \", \")", lk)
	f.W(")", -1)
	f.LB()

	f.W("type "+lk+"DTO struct {", 1)
	dtoMsg := "%-" + fmt.Sprintf("%d", maxKeyLength) + "s %-" + fmt.Sprintf("%d", maxDTOTypeLength) + "s %s"
	dtoFieldMsg := "%-" + fmt.Sprintf("%d", maxKeyLength+1) + "s %s.%s,"
	for _, fld := range m.Fields {
		typ, imps := typeString(fld.Type, fm, "dto")
		for _, imp := range imps {
			f.AddImport(imp.String())
		}
		f.Wf(dtoMsg, util.ToCamel(fld.Key), typ, "`db:\""+fld.Key+"\"`")
	}
	f.W("}", -1)
	f.LB()
	f.W(fmt.Sprintf("func (%s *%sDTO) To%s() *%s {", firstChar, lk, sk, sk), 1)
	f.W("return &"+sk+"{", 1)
	for _, fld := range m.Fields {
		call := util.ToCamel(fld.Key)
		switch typ, _ := typeString(fld.Type, fm, "dto"); typ {
		case "sql.NullBool":
			call += ".Bool"
		case "sql.NullString":
			call += ".String"
		}
		f.Wf(dtoFieldMsg, util.ToCamel(fld.Key)+":", firstChar, call)
	}
	f.W("}", -1)
	f.W("}", -1)
	f.LB()
	f.Wf("type %sDTOs []*%sDTO", lk, lk)
	f.LB()

	f.W(fmt.Sprintf("func (%s %sDTOs) To%s() %s {", firstChar, lk, pluralk, pluralk), 1)
	f.Wf("ret := make(%s, 0, len(%s))", pluralk, firstChar)
	f.W(fmt.Sprintf("for _, dto := range %s {", firstChar), 1)
	f.Wf("ret = append(ret, dto.To%s())", sk)
	f.W("}", -1)
	f.W("return ret")
	f.W("}", -1)

	ret := &Result{Key: "model", Out: f}
	return ret, nil
}

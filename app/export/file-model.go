package export

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

// nolint
func fileModel(m *model.Model, logger *zap.SugaredLogger) (*Result, error) {
	f := NewGoFile(m.Pkg, m.Key)

	pk := m.GetPK(logger)

	lk := util.StringToLowerCamel(util.StringToSingular(m.Key))
	sk := util.StringToCamel(util.StringToSingular(m.Key))
	pluralk := util.StringToCamel(util.StringToPlural(m.Key))
	firstChar := strings.ToLower(string(m.Key[0]))

	maxKeyLength := 0
	maxTypeLength := 0
	maxDTOTypeLength := 0
	dataFields := make([]string, 0, len(m.Fields))
	for _, fld := range m.Fields {
		if len(fld.Key) > maxKeyLength {
			maxKeyLength = len(fld.Key)
		}
		x, _ := typeString(fld.Type, "model")
		if len(x) > maxTypeLength {
			maxTypeLength = len(x)
		}
		z, _ := typeString(fld.Type, "dto")
		if len(z) > maxDTOTypeLength {
			maxDTOTypeLength = len(z)
		}
		dataFields = append(dataFields, firstChar+"."+util.StringToCamel(fld.Key))
	}

	f.W("type "+sk+" struct {", 1)
	msg := "%-" + fmt.Sprintf("%d", maxKeyLength) + "s %-" + fmt.Sprintf("%d", maxTypeLength) + "s %s%s"
	for _, fld := range m.Fields {
		typ, imps := typeString(fld.Type, "model")
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
		f.Wf(msg, util.StringToCamel(fld.Key), typ, "`json:\""+util.StringToLowerCamel(fld.Key)+omit+"\"`", suffix)
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

	f.W("type dto struct {", 1)
	dtoMsg := "%-" + fmt.Sprintf("%d", maxKeyLength) + "s %-" + fmt.Sprintf("%d", maxDTOTypeLength) + "s %s"
	dtoFieldMsg := "%-" + fmt.Sprintf("%d", maxKeyLength+1) + "s d.%s,"
	for _, fld := range m.Fields {
		typ, imps := typeString(fld.Type, "dto")
		for _, imp := range imps {
			f.AddImport(imp.String())
		}
		f.Wf(dtoMsg, util.StringToCamel(fld.Key), typ, "`db:\""+fld.Key+"\"`")
	}
	f.W("}", -1)
	f.LB()
	f.W(fmt.Sprintf("func (d *dto) To%s() *%s {", sk, sk), 1)
	f.W("return &"+sk+"{", 1)
	for _, fld := range m.Fields {
		call := util.StringToCamel(fld.Key)
		switch typ, _ := typeString(fld.Type, "dto"); typ {
		case "sql.NullBool":
			call += ".Bool"
		case "sql.NullString":
			call += ".String"
		}
		f.Wf(dtoFieldMsg, util.StringToCamel(fld.Key)+":", call)
	}
	f.W("}", -1)
	f.W("}", -1)
	f.LB()
	f.Wf("type dtos []*dto")
	f.LB()

	f.W(fmt.Sprintf("func (d dtos) To%s() %s {", pluralk, pluralk), 1)
	f.Wf("ret := make(%s, 0, len(d))", pluralk)
	f.W("for _, x := range d {", 1)
	f.Wf("ret = append(ret, x.To%s())", sk)
	f.W("}", -1)
	f.W("return ret")
	f.W("}", -1)

	ret := &Result{Key: f.Path(), Out: f}
	return ret, nil
}

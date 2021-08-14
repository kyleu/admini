package export

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

func goModelFile(m *model.Model, fm *Format, logger *zap.SugaredLogger) *Result {
	f := NewGoFile(m.Pkg, m.Key)

	pk := m.GetPK(logger)
	maxKeyLength := 0
	maxTypeLength := 0
	maxDTOTypeLength := 0
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
	}

	lk := util.ToLowerCamel(util.ToSingular(m.Key))
	sk := util.ToCamel(util.ToSingular(m.Key))

	f.W("type "+sk+" struct {", 1)
	msg := "%-" + fmt.Sprintf("%d", maxKeyLength) + "s %-" + fmt.Sprintf("%d", maxTypeLength) + "s %s%s"
	for _, fld := range m.Fields {
		typ, imp := typeString(fld.Type, fm, "model")
		if len(imp) > 0 {
			logger.Warn("imports...")
		}
		omit := ""
		if fld.Type.IsOption() {
			omit = ",omitempty"
		}
		suffix := ""
		if util.StringArrayContains(pk, fld.Key) {
			suffix = " /* primary key */"
		}
		f.W(fmt.Sprintf(msg, util.ToCamel(fld.Key), typ, "`json:\""+util.ToLowerCamel(fld.Key)+omit+"\"`", suffix))
	}
	f.W("}", -1)
	f.W("")
	f.W(fmt.Sprintf("type %ss []*%s", sk, sk))

	f.W("")
	f.W("type "+lk+"DTO struct {", 1)
	dtoMsg := "%-" + fmt.Sprintf("%d", maxKeyLength) + "s %-" + fmt.Sprintf("%d", maxDTOTypeLength) + "s %s"
	dtoFieldMsg := "%-" + fmt.Sprintf("%d", maxKeyLength + 1) + "s %s.%s,"
	for _, fld := range m.Fields {
		typ, imp := typeString(fld.Type, fm, "dto")
		if len(imp) > 0 {
			logger.Warn("imports...")
		}
		f.W(fmt.Sprintf(dtoMsg, util.ToCamel(fld.Key), typ, "`db:\""+fld.Key+"\"`"))
	}
	f.W("}", -1)
	f.W("")
	firstChar := strings.ToLower(string(m.Key[0]))
	f.W(fmt.Sprintf("func (%s *%sDTO) To%s() *%s {", firstChar, lk, sk, sk), 1)
	f.W(fmt.Sprintf("return &%s{", sk), 1)
	for _, fld := range m.Fields {
		call := util.ToCamel(fld.Key)
		switch typ, _ := typeString(fld.Type, fm, "dto"); typ {
		case "sql.NullBool":
			call += ".Bool"
		case "sql.NullString":
			call += ".String"
		}
		f.W(fmt.Sprintf(dtoFieldMsg, util.ToCamel(fld.Key) + ":", firstChar, call))
	}
	f.W("}", -1)
	f.W("}", -1)
	f.W("")
	f.W(fmt.Sprintf("type %sDTOs []*%sDTO", lk, lk))

	ret := &Result{Key: "model", Out: f}
	return ret
}

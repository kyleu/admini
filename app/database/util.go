package database

import (
	"strings"

	"github.com/kyleu/admini/app/util"
)

func valueStrings(values []interface{}) string {
	ret := util.StringArrayFromInterfaces(values)
	return strings.Join(ret, ", ")
}

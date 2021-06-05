package database

import (
	"strings"

	"github.com/kyleu/admini/app/util"
)

// Converts provided array elements to strings, then joins them as a list
func valueStrings(values []interface{}) string {
	ret := util.StringArrayFromInterfaces(values)
	return strings.Join(ret, ", ")
}

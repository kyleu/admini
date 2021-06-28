package util

import (
	"fmt"
	"strings"
	"unicode"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func SplitString(s string, sep byte, cutc bool) (string, string) {
	i := strings.IndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}

func SplitStringLast(s string, sep byte, cutc bool) (string, string) {
	i := strings.LastIndexByte(s, sep)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+1:]
	}
	return s[:i], s[i:]
}

func SplitAndTrim(s string, delim string) []string {
	split := strings.Split(s, delim)
	ret := make([]string, 0, len(split))
	for _, x := range split {
		x = strings.TrimSpace(x)
		if len(x) > 0 {
			ret = append(ret, x)
		}
	}
	return ret
}

var acronyms = []string{"Id"}

func ToCamel(s string) string {
	return acr(strcase.ToCamel(s))
}

func ToTitle(s string) string {
	ret := strings.Builder{}
	runes := []rune(ToCamel(s))
	for idx, c := range runes {
		if idx > 0 && idx < len(runes) && unicode.IsUpper(c) {
			if !unicode.IsUpper(runes[idx+1]) {
				ret.WriteRune(' ')
			} else if !unicode.IsUpper(runes[idx-1]) {
				ret.WriteRune(' ')
			}
		}
		ret.WriteRune(c)
	}
	return ret.String()
}

func ToLowerCamel(s string) string {
	return acr(strcase.ToLowerCamel(s))
}

var plrl *pluralize.Client

func ToPlural(s string) string {
	if plrl == nil {
		plrl = pluralize.NewClient()
	}
	return plrl.Plural(s)
}

func ToSingular(s string) string {
	if plrl == nil {
		plrl = pluralize.NewClient()
	}
	return plrl.Singular(s)
}

func StringForms(s string) (string, string) {
	if plrl == nil {
		plrl = pluralize.NewClient()
	}
	if plrl.IsSingular(s) {
		return s, plrl.Plural(s)
	} else {
		return plrl.Singular(s), s
	}
}

func Plural(count int, s string) string {
	var x string
	if count == 1 {
		x = ToSingular(s)
	} else {
		x = ToPlural(s)
	}
	return fmt.Sprintf("%d %s", count, x)
}

func acr(ret string) string {
	for _, a := range acronyms {
		for {
			i := strings.Index(ret, a)
			if i == -1 {
				break
			}
			ret = ret[:i] + strings.ToUpper(a) + ret[i+len(a):]
		}
	}
	return ret
}

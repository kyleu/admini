package util

import (
	"fmt"
	"strings"
	"unicode"

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

func Plural(count int, sing string, plur string) string {
	x := sing
	if count != 1 {
		x = plur
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

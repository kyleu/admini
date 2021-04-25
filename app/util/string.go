package util

import (
	"strings"
)

// Splits a string on the first instance of a provided byte, returning strings representing each side
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

// Splits a string on the last instance of a provided byte, returning strings representing each side
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

// Splits a string according to a delimeter, then trims each entry, filtering empty strings
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

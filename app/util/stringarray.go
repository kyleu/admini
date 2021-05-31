package util

import "fmt"

func StringArrayCopy(a []string) []string {
	ret := make([]string, 0, len(a))
	return append(ret, a...)
}

func StringArrayFromInterfaces(a []interface{}) []string {
	ret := make([]string, 0, len(a))
	for _, x := range a {
		ret = append(ret, fmt.Sprint(x))
	}
	return ret
}

func StringArrayContains(a []string, s string) bool {
	return StringArrayIndex(a, s) > -1
}

func StringArrayIndex(a []string, s string) int {
	for i, x := range a {
		if x == s {
			return i
		}
	}
	return -1
}

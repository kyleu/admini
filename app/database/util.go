package database

import (
	"fmt"
	"strings"
)

// Converts a string array into a SQL array string
func ArrayToString(a []string) string {
	return "{" + strings.Join(a, ",") + "}"
}

// Formats a SQL array string into a string array
func StringToArray(s string) []string {
	split := strings.Split(strings.TrimPrefix(strings.TrimSuffix(s, "}"), "{"), ",")
	ret := make([]string, 0)

	for _, x := range split {
		y := strings.TrimSpace(x)
		if len(y) > 0 {
			ret = append(ret, y)
		}
	}

	return ret
}

// Converts provided array elements to strings, then joins them as a list
func valueStrings(values []interface{}) string {
	ret := make([]string, 0, len(values))
	for _, v := range values {
		ret = append(ret, fmt.Sprintf(`"%v"`, v))
	}
	return strings.Join(ret, ", ")
}
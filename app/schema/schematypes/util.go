package schematypes

import "fmt"

func invalidInput(key string, v interface{}) string {
	return fmt.Sprintf("unable to parse [%v] from [%v] (%T)", key, v, v)
}

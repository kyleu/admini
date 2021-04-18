package schematypes

import (
	"fmt"
)

type Type interface {
	Key() string
	fmt.Stringer
}

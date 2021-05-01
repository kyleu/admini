package schematypes

import (
	"fmt"
)

type Type interface {
	Key() string
	Sortable() bool
	fmt.Stringer
}

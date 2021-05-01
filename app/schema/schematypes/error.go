package schematypes

const KeyError = "error"

type Error struct {
	Message string `json:"message"`
}

var _ Type = (*Error)(nil)

func (t *Error) Key() string {
	return KeyError
}

func (t *Error) String() string {
	return "error(" + t.Message + ")"
}

func (t *Error) Sortable() bool {
	return false
}

func NewError(msg string) *Wrapped {
	return Wrap(&Error{Message: msg})
}

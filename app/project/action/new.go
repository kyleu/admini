package action

import (
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func NewAction(args []string, typ Type, pkg util.Pkg) (*Action, error) {
	switch typ {
	case ActionTypeAll:
		return &Action{
			Key:         "TODO",
			Type:        ActionTypeAll.Key,
			Title:       "All Sources",
			Pkg:         pkg,
		}, nil
	default:
		return nil, errors.New("can't create unhandled action [" + typ.Key + "]")
	}
}


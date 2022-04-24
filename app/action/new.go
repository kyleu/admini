package action

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app/util"
)

func newAction(args []string, title string, typ *Type, pkg util.Pkg) (*Action, error) {
	if typ == nil {
		return nil, errors.New("action has nil type")
	}
	cfg := util.ValueMap{}
	switch typ {
	case TypeFolder:
		return base(typ, "folder", title, pkg, cfg), nil
	case TypeStatic:
		return base(typ, typ.Key, title, pkg, cfg), nil
	case TypeSeparator:
		return base(typ, "", "", pkg, cfg), nil
	case TypeAll:
		return base(typ, typ.Key, title, pkg, cfg), nil
	case TypeSource:
		if len(args) == 0 {
			return nil, errors.New("require one argument")
		}
		srcKey := args[0]
		cfg[TypeSource.Key] = srcKey
		return base(typ, srcKey, title, pkg, cfg), nil
	case TypePackage:
		if len(args) < 2 {
			return nil, errors.Errorf("require at least two arguments, observed [%d]", len(args))
		}
		srcKey := args[0]
		cfg[TypeSource.Key] = srcKey
		key := args[len(args)-1]
		cfg[TypePackage.Key] = strings.Join(args[1:], "/")
		return base(typ, key, title, pkg, cfg), nil
	case TypeModel:
		if len(args) < 2 {
			return nil, errors.Errorf("require at least two arguments, observed [%d]", len(args))
		}
		srcKey := args[0]
		cfg[TypeSource.Key] = srcKey
		key := args[len(args)-1]
		cfg[TypeModel.Key] = strings.Join(args[1:], "/")
		return base(typ, key, title, pkg, cfg), nil
	case TypeActivity:
		if len(args) != 2 {
			return nil, errors.Errorf("require exactly two arguments, observed [%d]", len(args))
		}
		srcKey := args[0]
		activity := args[1]
		cfg[TypeSource.Key] = srcKey
		cfg[TypeActivity.Key] = activity
		return base(typ, fmt.Sprintf("%s-%s", srcKey, activity), title, pkg, cfg), nil
	default:
		return nil, errors.Errorf("can't create unhandled action [%s]", typ.Key)
	}
}

func base(typ *Type, key string, title string, pkg util.Pkg, cfg util.ValueMap) *Action {
	if cfg == nil {
		cfg = util.ValueMap{}
	}
	return &Action{Key: "__" + key, TypeKey: typ.Key, Title: title, Pkg: pkg, Config: cfg}
}

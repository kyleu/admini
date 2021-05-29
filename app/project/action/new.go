package action

import (
	"strings"

	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func NewAction(args []string, typ Type, pkg util.Pkg) (*Action, error) {
	cfg := map[string]string{}
	switch typ {
	case TypeFolder:
		return base(typ, "folder", "New Folder", pkg, cfg), nil
	case TypeAll:
		return base(typ, typ.Key, "All Sources", pkg, cfg), nil
	case TypeSource:
		if len(args) == 0 {
			return nil, errors.New("require one argument")
		}
		srcKey := args[0]
		cfg["source"] = srcKey
		return base(typ, srcKey, srcKey, pkg, cfg), nil
	case TypeModel:
		if len(args) < 2 {
			return nil, errors.Errorf("require at least two arguments, observed [%v]", len(args))
		}
		srcKey := args[0]
		cfg["source"] = srcKey
		key := args[len(args)-1]
		cfg["model"] = strings.Join(args[1:], "/")
		return base(typ, key, key, pkg, cfg), nil
	case TypePackage:
		if len(args) < 2 {
			return nil, errors.Errorf("require at least two arguments, observed [%v]", len(args))
		}
		srcKey := args[0]
		cfg["source"] = srcKey
		key := args[len(args)-1]
		cfg["package"] = strings.Join(args[1:], "/")
		return base(typ, key, key, pkg, cfg), nil
	case TypeActivity:
		if len(args) != 2 {
			return nil, errors.Errorf("require exactly two arguments, observed [%v]", len(args))
		}
		srcKey := args[0]
		activity := args[1]
		cfg["source"] = srcKey
		cfg["activity"] = activity
		return base(typ, srcKey+"-"+activity, srcKey+" SQL", pkg, cfg), nil
	default:
		return nil, errors.New("can't create unhandled action [" + typ.Key + "]")
	}
}

func base(typ Type, key string, title string, pkg util.Pkg, cfg map[string]string) *Action {
	return &Action{Key: "__" + key, Type: typ.Key, Title: title, Pkg: pkg, Config: cfg}
}

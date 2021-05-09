package action

import (
	"fmt"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"path/filepath"
)

func Load(root string, files filesystem.FileLoader) (Actions, error) {
	return loadChildren(root, files, util.Pkg{})
}

func loadChildren(dir string, files filesystem.FileLoader, pkg util.Pkg) (Actions, error) {
	ret := Actions{}
	kids := files.ListDirectories(dir)
	ret = make(Actions, 0, len(kids))
	for _, kid := range kids {
		p := append(pkg, kid)
		x, err := loadAction(dir, kid, files, p)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error loading [%v]", kid))
		}
		ret = append(ret, x)
	}

	ret.Sort()

	return ret, nil
}

func loadAction(dir string, key string, files filesystem.FileLoader, pkg util.Pkg) (*Action, error) {
	kp := filepath.Join(dir, key)
	js := filepath.Join(kp, "action.json")
	ret := &Action{}

	if files.Exists(js) {
		out, err := files.ReadFile(js)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read action ["+kp+"]")
		}

		err = util.FromJSON(out, ret)
		if err != nil {
			return nil, errors.Wrap(err, "unable to parse action")
		}
		ret.Key = key
		ret.Pkg = pkg
	}

	x, err := loadChildren(kp, files, pkg)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load action children")
	}

	ret.Children = x

	return ret, nil
}

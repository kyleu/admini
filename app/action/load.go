package action

import (
	"path/filepath"
	"strings"

	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func Load(root string, files filesystem.FileLoader) (Actions, error) {
	return loadChildren(root, files, util.Pkg{})
}

func loadChildren(dir string, files filesystem.FileLoader, pkg util.Pkg) (Actions, error) {
	kids := files.ListDirectories(dir)
	ret := make(Actions, 0, len(kids))
	for _, kid := range kids {
		if strings.HasPrefix(kid, ".") {
			continue
		}
		x, err := loadAction(dir, kid, files, pkg)
		if err != nil {
			return nil, errors.Wrapf(err, "error loading [%s]", kid)
		}
		ret = append(ret, x)
	}

	ret.Sort()

	return ret, nil
}

func loadAction(dir string, key string, files filesystem.FileLoader, pkg util.Pkg) (*Action, error) {
	kp := filepath.Join(dir, key)
	js := filepath.Join(kp, "action.json")
	d := &dto{}

	if files.Exists(js) {
		out, err := files.ReadFile(js)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read action [%s]", kp)
		}

		err = util.FromJSON(out, d)
		if err != nil {
			return nil, errors.Wrap(err, "unable to parse action")
		}
	}

	d.Pkg = pkg
	ret := d.ToAction(key)
	x, err := loadChildren(kp, files, append(pkg, key))
	if err != nil {
		return nil, errors.Wrap(err, "unable to load action children")
	}
	ret.Children = x
	return ret, nil
}

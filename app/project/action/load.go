package action

import (
	"fmt"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
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
			return nil, fmt.Errorf("error loading [%v]: %w", kid, err)
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
			return nil, fmt.Errorf("unable to read action ["+kp+"]: %w", err)
		}

		err = util.FromJSON(out, ret)
		if err != nil {
			return nil, fmt.Errorf("unable to parse action: %w", err)
		}
		ret.Key = key
		ret.Pkg = pkg
	}

	x, err := loadChildren(kp, files, pkg)
	if err != nil {
		return nil, fmt.Errorf("unable to load action children: %w", err)
	}

	ret.Children = x

	return ret, nil
}

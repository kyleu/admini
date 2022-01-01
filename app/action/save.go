package action

import (
	"path/filepath"

	"github.com/kyleu/admini/app/lib/filesystem"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func SaveAll(prj string, acts Actions, files filesystem.FileLoader) (int, error) {
	prjPath := filepath.Join("project", prj)
	if !files.Exists(prjPath) {
		return 0, errors.Errorf("project directory [%s] does not exist", prjPath)
	}
	actPath := filepath.Join(prjPath, "actions~")
	if files.Exists(actPath) {
		_ = files.RemoveRecursive(actPath)
	}

	err := files.CreateDirectory(actPath)
	if err != nil {
		return 0, errors.Wrapf(err, "can't create actions directory at [%s]", actPath)
	}

	count := 0
	for _, act := range acts {
		c, e := Save(actPath, act, files)
		if e != nil {
			return 0, e
		}
		count += c
	}

	err = replace(prjPath, "actions~", "actions", files)
	if err != nil {
		return 0, errors.Wrap(err, "can't replace actions directory")
	}

	return count, nil
}

func Save(path string, act *Action, files filesystem.FileLoader) (int, error) {
	count := 1
	dest := filepath.Join(path, act.Key)
	if !files.Exists(dest) {
		err := files.CreateDirectory(dest)
		if err != nil {
			return 0, errors.Wrapf(err, "unable to create directory for action [%s]", act.Key)
		}
	}
	actFile := filepath.Join(dest, "action.json")
	dto := newDTO(act)
	js := util.ToJSONBytes(dto, true)
	if len(js) != 2 {
		err := files.WriteFile(actFile, js, filesystem.DefaultMode, true)
		if err != nil {
			return 0, errors.Wrapf(err, "unable to write file for action [%s]", act.Key)
		}
	}

	for _, kid := range act.Children {
		c, err := Save(dest, kid, files)
		if err != nil {
			return 0, err
		}
		count += c
	}

	return count, nil
}

func replace(root string, src string, tgt string, files filesystem.FileLoader) error {
	srcPath := filepath.Join(root, src)
	tgtPath := filepath.Join(root, tgt)
	tmpPath := filepath.Join(root, "~tmp")
	tgtPathExists := files.Exists(tgtPath)

	if tgtPathExists {
		err := files.Move(tgtPath, tmpPath)
		if err != nil {
			return err
		}
	}

	err := files.Move(srcPath, tgtPath)
	if err != nil {
		return err
	}

	if tgtPathExists {
		err = files.RemoveRecursive(tmpPath)
		if err != nil {
			return err
		}
	}

	return nil
}

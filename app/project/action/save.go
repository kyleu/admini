package action

import (
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"path/filepath"
)

func Save(prj string, acts Actions, files filesystem.FileLoader) error {
	prjPath := filepath.Join("project", prj)
	if !files.Exists(prjPath) {
		return errors.New("project directory [" + prjPath + "] does not exist")
	}
	actPath := filepath.Join(prjPath, "actions~")
	if !files.Exists(actPath) {
		err := files.CreateDirectory(actPath)
		if err != nil {
			return errors.Wrap(err, "can't create actions directory at [" + actPath + "]")
		}
	}

	for _, act := range acts {
		err := saveAction(actPath, act, files)
		if err != nil {
			return err
		}
	}

	err := replace(prjPath, "actions~", "actions", files)
	if err != nil {
		return errors.Wrap(err, "can't replace actions directory")
	}

	return nil
}

func replace(root string, src string, tgt string, files filesystem.FileLoader) error {
	srcPath := filepath.Join(root, src)
	tgtPath := filepath.Join(root, tgt)
	tmpPath := filepath.Join(root, "~tmp")

	err := files.Move(tgtPath, tmpPath)
	if err != nil {
		return err
	}

	err = files.Move(srcPath, tgtPath)
	if err != nil {
		return err
	}

	err = files.RemoveRecursive(tmpPath)
	if err != nil {
		return err
	}

	return nil
}

func saveAction(path string, act *Action, files filesystem.FileLoader) error {
	dest := filepath.Join(path, act.Key)
	if !files.Exists(dest) {
		err := files.CreateDirectory(dest)
		if err != nil {
			return errors.Wrap(err, "unable to create directory for action [" + act.Key + "]")
		}
	}
	actFile := filepath.Join(dest, "action.json")
	dto := newDTO(act)
	js := util.ToJSONBytes(dto, true)
	if len(js) != 2 {
		err := files.WriteFile(actFile, js, true)
		if err != nil {
			return errors.Wrap(err, "unable to write file for action ["+act.Key+"]")
		}
	}

	for _, kid := range act.Children {
		err := saveAction(dest, kid, files)
		if err != nil {
			return err
		}
	}

	return nil
}

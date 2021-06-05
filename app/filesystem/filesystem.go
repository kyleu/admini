package filesystem

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

type FileSystem struct {
	root   string
	logger *zap.SugaredLogger
}

var _ FileLoader = (*FileSystem)(nil)

func NewFileSystem(root string, logger *zap.SugaredLogger) *FileSystem {
	return &FileSystem{root: root, logger: logger.With(zap.String("service", "filesystem"))}
}

func (f *FileSystem) getPath(ss ...string) string {
	s := filepath.Join(ss...)
	if strings.HasPrefix(s, f.root) {
		return s
	}
	return filepath.Join(f.root, s)
}

func (f *FileSystem) Root() string {
	return f.root
}

func (f *FileSystem) ReadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(f.getPath(path))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read file [%s]", path)
	}
	return b, nil
}

func (f *FileSystem) CreateDirectory(path string) error {
	p := f.getPath(path)
	if err := os.MkdirAll(p, 0o755); err != nil {
		return errors.Wrapf(err, "unable to create data directory [%s]", p)
	}
	return nil
}

func (f *FileSystem) WriteFile(path string, content []byte, overwrite bool) error {
	p := f.getPath(path)
	_, err := os.Stat(p)
	if os.IsExist(err) && !overwrite {
		return errors.Errorf("file [%s] exists, will not overwrite", p)
	}
	dd := filepath.Dir(path)
	err = f.CreateDirectory(dd)
	if err != nil {
		return errors.Wrapf(err, "unable to create data directory [%s]", dd)
	}
	file, err := os.Create(p)
	if err != nil {
		return errors.Wrapf(err, "unable to create file [%s]", p)
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write(content)
	if err != nil {
		return errors.Wrapf(err, "unable to write content to file [%s]", p)
	}
	return nil
}

func (f *FileSystem) CopyFile(src string, tgt string) error {
	sp := f.getPath(src)
	tp := f.getPath(tgt)

	if targetExists := f.Exists(tp); targetExists {
		return errors.Errorf("file [%s] exists, will not overwrite", tp)
	}

	input, err := f.ReadFile(sp)
	if err != nil {
		return err
	}

	err = f.WriteFile(tp, input, false)
	return err
}

func (f *FileSystem) Move(src string, tgt string) error {
	sp := f.getPath(src)
	if sourceExists := f.Exists(sp); !sourceExists {
		return errors.Errorf("source file [%s] does not exist, can't move", sp)
	}

	tp := f.getPath(tgt)
	if targetExists := f.Exists(tp); targetExists {
		return errors.Errorf("target file [%s] exists, will not overwrite", tp)
	}

	if err := os.Rename(sp, tp); err != nil {
		return errors.Wrapf(err, "error renaming [%s] to [%s]", sp, tp)
	}

	return nil
}

func (f *FileSystem) ListJSON(path string) []string {
	return f.ListExtension(path, "json")
}

func (f *FileSystem) ListExtension(path string, ext string) []string {
	glob := "*." + ext
	matches, err := filepath.Glob(f.getPath(path, glob))
	if err != nil {
		f.logger.Warnf("cannot list [%s] in path [%s]: %+v", ext, path, err)
	}
	ret := make([]string, 0, len(matches))
	for _, j := range matches {
		idx := strings.LastIndex(j, "/")
		if idx > 0 {
			j = j[idx+1:]
		}
		ret = append(ret, strings.TrimSuffix(j, "."+ext))
	}
	return ret
}

func (f *FileSystem) ListDirectories(path string) []string {
	if !f.Exists(path) {
		return nil
	}
	p := f.getPath(path)
	files, err := ioutil.ReadDir(p)
	if err != nil {
		f.logger.Warnf("cannot list path [%s]: %+v", path, err)
	}
	var ret []string
	for _, f := range files {
		if f.IsDir() {
			ret = append(ret, f.Name())
		}
	}
	return ret
}

func (f *FileSystem) Exists(path string) bool {
	p := f.getPath(path)
	_, err := os.Stat(p)
	return err == nil
}

func (f *FileSystem) IsDir(path string) bool {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err == nil {
		return s.IsDir()
	}
	return false
}

func (f *FileSystem) Remove(path string) error {
	p := f.getPath(path)
	f.logger.Warnf("removing file at path [%s]", p)
	if err := os.Remove(p); err != nil {
		return errors.Wrapf(err, "error removing file [%s]", path)
	}
	return nil
}

func (f *FileSystem) RemoveRecursive(path string) error {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err != nil {
		return errors.Wrapf(err, "unable to stat file [%s]", path)
	}
	if s.IsDir() {
		var files []fs.FileInfo
		files, err = ioutil.ReadDir(p)
		if err != nil {
			f.logger.Warnf("cannot read path [%s] for removal: %+v", path, err)
		}
		for _, file := range files {
			err = f.RemoveRecursive(filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	err = os.Remove(p)
	if err != nil {
		return errors.Wrapf(err, "unable to remove file [%s]", path)
	}
	return nil
}

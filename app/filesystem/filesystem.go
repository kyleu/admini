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

// Constructor
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

// Root directory, as a string
func (f *FileSystem) Root() string {
	return f.root
}

// Reads the contents of a file as a byte array
func (f *FileSystem) ReadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(f.getPath(path))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read file [%v]", path)
	}
	return b, nil
}

// Creates a directory, like it says on the tin
func (f *FileSystem) CreateDirectory(path string) error {
	p := f.getPath(path)
	if err := os.MkdirAll(p, 0o755); err != nil {
		return errors.Wrapf(err, "unable to create data directory [%v]", p)
	}
	return nil
}

// Writes the the provided byte array to a file
func (f *FileSystem) WriteFile(path string, content []byte, overwrite bool) error {
	p := f.getPath(path)
	_, err := os.Stat(p)
	if os.IsExist(err) && !overwrite {
		return errors.New("file [" + p + "] exists, will not overwrite")
	}
	dd := filepath.Dir(path)
	err = f.CreateDirectory(dd)
	if err != nil {
		return errors.Wrapf(err, "unable to create data directory [%v]", dd)
	}
	file, err := os.Create(p)
	if err != nil {
		return errors.Wrapf(err, "unable to create file [%v]", p)
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write(content)
	if err != nil {
		return errors.Wrapf(err, "unable to write content to file [%v]", p)
	}
	return nil
}

// Copies the contents of one file to another
func (f *FileSystem) CopyFile(src string, tgt string) error {
	sp := f.getPath(src)
	tp := f.getPath(tgt)

	if targetExists := f.Exists(tp); targetExists {
		return errors.New("file [" + tp + "] exists, will not overwrite")
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
		return errors.New("source file [" + sp + "] does not exist, can't move")
	}

	tp := f.getPath(tgt)
	if targetExists := f.Exists(tp); targetExists {
		return errors.New("target file [" + tp + "] exists, will not overwrite")
	}

	if err := os.Rename(sp, tp); err != nil {
		return errors.Wrapf(err, "error renaming [%v] to [%v]", sp, tp)
	}

	return nil
}

// Lists all files in a directory with a `.json` extension
func (f *FileSystem) ListJSON(path string) []string {
	return f.ListExtension(path, "json")
}

// Lists all files in a directory with a provided extension
func (f *FileSystem) ListExtension(path string, ext string) []string {
	glob := "*." + ext
	matches, err := filepath.Glob(f.getPath(path, glob))
	if err != nil {
		f.logger.Warnf("cannot list [%v] in path [%v]: %+v", ext, path, err)
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

// Lists all directories in a directory
func (f *FileSystem) ListDirectories(path string) []string {
	if !f.Exists(path) {
		return nil
	}
	p := f.getPath(path)
	files, err := ioutil.ReadDir(p)
	if err != nil {
		f.logger.Warnf("cannot list path [%v]: %+v", path, err)
	}
	var ret []string
	for _, f := range files {
		if f.IsDir() {
			ret = append(ret, f.Name())
		}
	}
	return ret
}

// Returns a boolean indicating if the file exists
func (f *FileSystem) Exists(path string) bool {
	p := f.getPath(path)
	_, err := os.Stat(p)
	return err == nil
}

// Returns a boolean indicating if the file exists and is a directory
func (f *FileSystem) IsDir(path string) bool {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err == nil {
		return s.IsDir()
	}
	return false
}

// Removes the file at the provided path
func (f *FileSystem) Remove(path string) error {
	p := f.getPath(path)
	f.logger.Warn("removing file at path [" + p + "]")
	if err := os.Remove(p); err != nil {
		return errors.Wrapf(err, "error removing file [%v]", path)
	}
	return nil
}

// Removes the file at the provided path, recursively
func (f *FileSystem) RemoveRecursive(path string) error {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err != nil {
		return errors.Wrapf(err, "unable to stat file [%v]", path)
	}
	if s.IsDir() {
		var files []fs.FileInfo
		files, err = ioutil.ReadDir(p)
		if err != nil {
			f.logger.Warnf("cannot read path [%v] for removal: %+v", path, err)
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
		return errors.Wrapf(err, "unable to remove file [%v]", path)
	}
	return nil
}

package filesystem

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kyleu/admini/app/util"
)

type FileSystem struct {
	root string
}

var _ FileLoader = (*FileSystem)(nil)

// Constructor
func NewFileSystem(root string) *FileSystem {
	return &FileSystem{root: root}
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
		return nil, fmt.Errorf("unable to read file [%v]: %w", path, err)
	}
	return b, nil
}

// Creates a directory, like it says on the tin
func (f *FileSystem) CreateDirectory(path string) error {
	p := f.getPath(path)
	if err := os.MkdirAll(p, 0o755); err != nil {
		return fmt.Errorf("unable to create data directory [%v]: %w", p, err)
	}
	return nil
}

// Writes the the provided byte array to a file
func (f *FileSystem) WriteFile(path string, content []byte, overwrite bool) error {
	p := f.getPath(path)
	_, err := os.Stat(p)
	if os.IsExist(err) && !overwrite {
		return fmt.Errorf("file [" + p + "] exists, will not overwrite")
	}
	dd := filepath.Dir(path)
	err = f.CreateDirectory(dd)
	if err != nil {
		return fmt.Errorf("unable to create data directory [%v]: %w", dd, err)
	}
	file, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("unable to create file [%v]: %w", p, err)
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("unable to write content to file [%v]: %w", p, err)
	}
	return nil
}

// Copies the contents of one file to another
func (f *FileSystem) CopyFile(src string, tgt string) error {
	sp := f.getPath(src)
	tp := f.getPath(tgt)

	if targetExists := f.Exists(tp); targetExists {
		return fmt.Errorf("file [" + tp + "] exists, will not overwrite")
	}

	input, err := f.ReadFile(sp)
	if err != nil {
		return err
	}

	err = f.WriteFile(tp, input, false)
	return err
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
		util.LogWarn("cannot list [%v] in path [%v]: %+v", ext, path, err)
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
	p := f.getPath(path)
	files, err := ioutil.ReadDir(p)
	if err != nil {
		util.LogWarn("cannot list path [%v]: %+v", path, err)
	}
	ret := make([]string, 0)
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
	util.LogWarn("removing file at path [" + p + "]")
	if err := os.Remove(p); err != nil {
		return fmt.Errorf("error removing file [%v]: %w", path, err)
	}
	return nil
}

// Removes the file at the provided path, recursively
func (f *FileSystem) RemoveRecursive(path string) error {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err != nil {
		return fmt.Errorf("unable to stat file [%v]: %w", path, err)
	}
	if s.IsDir() {
		var files []fs.FileInfo
		files, err = ioutil.ReadDir(p)
		if err != nil {
			util.LogWarn("cannot list path ["+path+"]: %+v", err)
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
		return fmt.Errorf("unable to remove file [%v]: %w", path, err)
	}
	return nil
}

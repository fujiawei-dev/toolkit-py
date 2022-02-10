{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	Home     = "~"
	HomePath = string('~' + filepath.Separator)
)

func Join(elem ...string) string {
	return filepath.Join(elem...)
}

func Rel(basePath, targetPath string) (string, error) {
	if !IsAbs(targetPath) {
		return targetPath, nil
	}
	return filepath.Rel(basePath, targetPath)
}

func MustRel(basePath, targetPath string) string {
	rel, _ := Rel(basePath, targetPath)
	return rel
}

func ToSlash(path string) string {
	return filepath.ToSlash(path)
}

func Abs(path string) (string, error) {
	if len(path) > 2 && path[:2] == HomePath {
		if usr, err := user.Current(); err == nil {
			path = Join(usr.HomeDir, path[2:])
		}
	}
	return filepath.Abs(path)
}

func MustAbs(path string) string {
	abs, _ := Abs(path)
	return abs
}

func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir returns if a path exists, and is a directory or symlink.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := info.Mode()
	return mode&os.ModeDir != 0 || mode&os.ModeSymlink != 0
}

// IsFile returns true if file exists and is not a directory.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func IsEmptyDir(path string) bool {
	items, err := ioutil.ReadDir(path)
	return len(items) == 0 && err == nil
}

func DeleteEmptyDir(path string) error {
	if IsEmptyDir(path) {
		return os.Remove(path)
	}
	return nil
}

func DeleteEmptyDirRecursive(path string) error {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if len(infos) == 0 {
		return os.Remove(path)
	}
	for _, info := range infos {
		if info.IsDir() {
			if err = DeleteEmptyDirRecursive(
				Join(path, info.Name()),
			); err != nil {
				return err
			}
		}
	}
	return DeleteEmptyDir(path)
}

func ListDir(path string) (files []string) {
	file, err := os.Open(path)
	if err != nil {
		return files
	}

	files, _ = file.Readdirnames(-1)
	return files
}

func ListDir2(path string) (files []string) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return files
	}
	for _, info := range infos {
		files = append(files, info.Name())
	}
	return files
}

func Ext(path string) string {
	return filepath.Ext(path)
}

func Base(path string) string {
	return filepath.Base(path)
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func IndexPathSeparator(path string) (index int) {
	for index < len(path) {
		if os.IsPathSeparator(path[index]) {
			return index
		}

		index++
	}

	return index
}

func RemovePathSeparator(path string) string {
	if os.IsPathSeparator(path[0]) {
		path = path[1:]
	}
	if os.IsPathSeparator(path[len(path)-1]) {
		path = path[:len(path)-1]
	}
	return path
}

// Split c:\images\cover.jpg => [c:/image, cover.jgp]
func Split(path string) (string, string) {
	dir, file := filepath.Split(path)

	return strings.TrimRight(dir, `/\`), file
}

func SplitPath(path string) (string, string, string) {
	if path == "" {
		return ".", "", ""
	}
	i, j := len(path)-1, len(path)
	if path[i] == ':' || path[i] == '/' {
		return path, "", ""
	}
	for i >= 0 && !os.IsPathSeparator(path[i]) {
		if path[i] == '.' && j == len(path) {
			j = i
		}
		i--
	}
	if i == 0 && j == len(path) {
		return "", path, ""
	}
	return path[:i+1], path[i+1 : j], path[j:]
}

func SplitPathExt(path string) (string, string) {
	ext := Ext(path)
	return path[:len(path)-len(ext)], ext
}

func GetFileNamePath(path string) string {
	fileNamePath, _ := SplitPathExt(path)
	return fileNamePath
}

func GetFileNameExt(path string) string {
	return Base(path)
}

func GetFileName(path string) string {
	_, fileName, _ := SplitPath(path)
	return fileName
}

func GetParentPath(path string) string {
	return Dir(path)
}

func ReplaceExt(path, ext string) string {
	return GetFileNamePath(path) + ext
}

func ReplaceFileName(path, fileName string) string {
	parent, _, ext := SplitPath(path)
	return Join(parent, fileName+ext)
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

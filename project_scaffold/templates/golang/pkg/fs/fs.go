{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"io/ioutil"
	"os"
)

// Overwrite overwrites the file with data. Creates file if not present.
func Overwrite(fileName string, data []byte) bool {
	f, err := os.Create(fileName)
	if err != nil {
		return false
	}

	_, err = f.Write(data)
	return err == nil
}

// PathWritable tests if a path exists and is writable.
func PathWritable(path string) bool {
	if !Exists(path) {
		return false
	}

	if f, err := ioutil.TempFile(path, ""); err != nil {
		return false
	} else if err = f.Close(); err != nil {
		return false
	} else if err = os.Remove(f.Name()); err != nil {
		return false
	}

	return true
}

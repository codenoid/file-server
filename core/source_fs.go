package core

import (
	"io/ioutil"
	"os"
	"path"
)

// ReadFile read local file by path
func ReadFile(folder, filename string) (val []byte, err error) {
	file, err := os.Open(path.Join(folder, filename))
	if err != nil {
		return val, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return val, err
	}

	return b, nil
}

package util

import (
	"fmt"
	"io"
	"os"
)

var File *fileUtil

type fileUtil struct{}

func init() {
	File = &fileUtil{}
}

// Read ...
func (f *fileUtil) Read(path string) ([]byte, error) {
	if exist, isFile := f.IsExist(path); !exist || !isFile {
		return nil, fmt.Errorf("文件不存在或者是一个目录")
	}

	fileHandler, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(fileHandler)
}

// IsExist ...
func (f *fileUtil) IsExist(filePath string) (bool, bool) {
	file, err := os.Stat(filePath)

	return nil == err || os.IsExist(err), (nil == err || os.IsExist(err)) && !file.IsDir()
}

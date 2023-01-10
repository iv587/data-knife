package file

import (
	"os"
	"path/filepath"
)

const rootPath = "resource"

// DbPath sqlite存储路径
var DbPath = filepath.Join(rootPath, "data", "data.sqlite")

// CreatFile 创建文件，并创建父目录
func CreatFile(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}
	_, err = os.Create(path)
	return err
}

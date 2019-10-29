package utils

import (
	"os"
	"path/filepath"
)

// GetPackageName 通过实际路径，得到包名
func GetPackageName(Path string) (string, error) {
	path, err := filepath.Abs(Path)
	if err == nil {
		var fileInfo os.FileInfo
		fileInfo, err = os.Stat(path)
		if err == nil {
			return fileInfo.Name(), nil
		}
	}
	return "", err
}

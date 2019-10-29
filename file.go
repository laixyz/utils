package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

/*
FileWrite 指定文件路径，创建或覆盖文件，并写入文件内容
*/
func FileWrite(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// FileHashByMD5 指定文件路径得到md5类型的hash值
func FileHashByMD5(path string) (hash string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(md5hash.Sum(nil)), nil
}

package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Create(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil || !os.IsNotExist(err) {
		return true
	}
	return false
}

func ReadBytes(filename string) ([]byte, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

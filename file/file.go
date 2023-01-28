package file

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/leigme/loki"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// CreateDir 创建文件夹
func CreateDir(filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(filename, os.ModePerm)
		}
		return err
	}
	if !fi.IsDir() {
		return errors.New(fmt.Sprintf("%s is exist file", filename))
	}
	return err
}

// Create 创建文件写入内容
func Create(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

// Md5 读取文件的md5值
func Md5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Delete 删除文件
func Delete(filename string) error {
	return os.Remove(filename)
}

// Exist 判断文件是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil || !os.IsNotExist(err) {
		return true
	}
	return false
}

// ReadBytes 读取文件的内容
func ReadBytes(filename string) ([]byte, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

// Split 根据块大小切分文件
func Split(filename string, chunkSize int64, execute func(num int, data []byte)) error {
	var (
		fi  os.FileInfo
		fr  *os.File
		err error
	)
	fi, err = os.Stat(filename)
	if err != nil {
		return err
	}
	if chunkSize <= 0 {
		return errors.New("chunk size must more than zero")
	}
	num := math.Ceil(float64(fi.Size()) / float64(chunkSize))
	fr, err = os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fr.Close()
	b := make([]byte, chunkSize)
	var i int64 = 1
	for ; i <= int64(num); i++ {
		fr.Seek((i-1)*chunkSize, 0)
		if len(b) > int(fi.Size()-(i-1)*chunkSize) {
			b = make([]byte, fi.Size()-(i-1)*chunkSize)
		}
		if _, err = fr.Read(b); err == nil {
			execute(int(i), b)
		}
	}
	return err
}

// deleteAllTmp 删除临时文件
func deleteAllTmp(dir, pre string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), pre) && strings.HasSuffix(f.Name(), loki.TmpSuffix) {
			err = Delete(filepath.Join(dir, f.Name()))
		}
	}
	return err
}

// Merge 合并临时文件
func Merge(dir, filename string) error {
	var (
		f    *os.File
		fi   os.FileInfo
		ts   ListWithName
		data []byte
		err  error
	)
	f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err = os.Stat(filename)
	if err != nil {
		return err
	}
	ts, err = os.ReadDir(dir)
	if err != nil {
		return err
	}
	sort.Stable(ts)
	fw := bufio.NewWriter(f)
	for _, t := range ts {
		if strings.HasPrefix(t.Name(), strings.TrimSuffix(fi.Name(), filepath.Ext(fi.Name()))) && strings.HasSuffix(t.Name(), loki.TmpSuffix) {
			if data, err = ReadBytes(filepath.Join(dir, t.Name())); err == nil {
				_, err = fw.Write(data)
				data = make([]byte, 0)
			}
		}
	}
	fw.Flush()
	if err == nil {
		return deleteAllTmp(filepath.Dir(filename), strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
	}
	return err
}

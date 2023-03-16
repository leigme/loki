package file

import (
	"bytes"
	"github.com/leigme/loki/param"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func CreateRequestBody(paramName, filename string) (string, io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	bw := multipart.NewWriter(buf)
	f, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()
	if md5, err := Md5(filename); err == nil {
		if pw, err := bw.CreateFormField(string(param.Md5)); err == nil {
			pw.Write([]byte(md5))
		}
	}
	if strings.EqualFold(paramName, "") {
		paramName = string(param.File)
	}
	fw, err := bw.CreateFormFile(paramName, filepath.Base(filename))
	if err != nil {
		return "", nil, err
	}
	_, err = io.Copy(fw, f)
	if err != nil {
		return "", nil, err
	}
	bw.Close()
	return bw.FormDataContentType(), buf, nil
}

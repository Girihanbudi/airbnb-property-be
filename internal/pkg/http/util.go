package http

import (
	"io/ioutil"
	"mime/multipart"
)

func ParseFileHeaderAsBytes(file multipart.FileHeader) (byteContainer []byte, err error) {
	var f multipart.File
	if f, err = file.Open(); err != nil {
		return
	}

	byteContainer, err = ioutil.ReadAll(f)
	return
}

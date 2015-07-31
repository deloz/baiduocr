package baiduocr

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	IMAGE_FIELD_NAME = "image"
	IMAGE_MAX_SIZE   = 307200 // 最大字节数, 300K字节
	IMAGE_EXT        = "jpg"  // 图片的格式, 只能是jpg
)

var (
	ErrImageInvalidSize = errors.New("图片过大, 最大只能为" + strconv.Itoa(IMAGE_MAX_SIZE) + "字节")
	ErrImageNotExist    = errors.New("请使用图片")
	ErrImageInvalid     = errors.New("图片格式不正确, 只能使用jpg扩展名")
)

func (o *Ocr) encodeImage() (string, error) {
	fileContents, err := ioutil.ReadFile(o.image)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(fileContents), nil
}

func (o *Ocr) detectImage(filename string) (bool, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return false, err
	}

	if fi.Size() > IMAGE_MAX_SIZE {
		return false, ErrImageInvalidSize
	}

	if fi.IsDir() {
		return false, ErrImageNotExist
	}

	ext := fi.Name()
	if !strings.HasSuffix(ext, IMAGE_EXT) {
		return false, ErrImageInvalid
	}

	return true, nil
}

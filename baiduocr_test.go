package baiduocr

import (
	"log"
	"os"
	"testing"
)

type TestOcr struct {
	*Ocr
}

func TestOcr_Scan(t *testing.T) {
	ocr := &Ocr{
		APIKey:       "your app key",
		FromDevice:   FROM_DEVICE_API,
		ClientIP:     "192.168.10.2",
		DetectType:   DETECT_TYPE_LOCATE_RECOGNIZE,
		LanguageType: LANGUAGE_TYPE_CHN_ENG,
		ImageType:    IMAGE_TYPE_ORIGINAL,
		//或者 使用base64编码
		//ImageType: IMAGE_TYPE_BASE64,
	}

	words, err := ocr.Scan("test.jpg")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	txt := ""
	for _, word := range words {
		txt += word
	}

	log.Println("--------")
	log.Println(txt)
	log.Println("--------")
}

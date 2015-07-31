package main

import (
	"log"
	"os"

	"github.com/deloz/baiduocr"
)

func main() {
	ocr := &baiduocr.Ocr{
		APIKey:       "your app key",
		FromDevice:   baiduocr.FROM_DEVICE_API,
		ClientIP:     "192.168.10.2",
		DetectType:   baiduocr.DETECT_TYPE_LOCATE_RECOGNIZE,
		LanguageType: baiduocr.LANGUAGE_TYPE_CHN_ENG,
		ImageType:    baiduocr.IMAGE_TYPE_ORIGINAL,
		//或者 使用base64编码
		//ImageType: baiduocr.IMAGE_TYPE_BASE64,
	}

	words, err := ocr.Scan("12306.jpg")
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

	words, err = ocr.Scan("jd.jpg")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	txt = ""
	for _, word := range words {
		txt += word
	}

	log.Println("--------")
	log.Println(txt)
	log.Println("--------")
}

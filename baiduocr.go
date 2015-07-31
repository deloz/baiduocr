package baiduocr

import (
	"errors"
	"strconv"
)

const (
	//客户端出口IP
	CLIENT_IP_DEFAULT = "10.10.10.0" // [默认]

	//来源
	FROM_DEVICE_API     = "go-baidu-ocr"
	FROM_DEVICE_ANDROID = "android"
	FROM_DEVICE_IPHONE  = "iPhone"
	FROM_DEVICE_PC      = "pic" // [默认]

	//OCR接口类型
	DETECT_TYPE_LOCATE_RECOGNIZE      = "LocateRecognize"     // LocateRecognize: 整图文字检测, 识别; 以行为单位 [默认]
	DETECT_TYPE_LOCATE                = "Locate"              // Locate: 整图文字行定位
	DETECT_TYPE_RECOGNIZE             = "Recognize"           // Recognize: 整图文字识别
	DETECT_TYPE_SINGLE_CHAR_RECOGNIZE = "SingleCharRecognize" // SingleCharRecognize: 单字图像识别

	// 图片资源类型
	IMAGE_TYPE_BASE64   = 1 // 数值1: 经过BASE64编码后的字串, 然后需要经过urlencode处理(特别重要) [默认]
	IMAGE_TYPE_ORIGINAL = 2 // 数值2: 图片原文件

	//要检测的文字类型
	LANGUAGE_TYPE_CHN_ENG = "CHN_ENG" // CHN_ENG(中英) [默认]
	LANGUAGE_TYPE_ENG     = "ENG"     // ENG(英文)
	LANGUAGE_TYPE_JAP     = "JAP"     // JAP(日文)
	LANGUAGE_TYPE_KOR     = "KOR"     // KOR(韩文)

	//接口地址
	API_URL = "http://apis.baidu.com/apistore/idlocr/ocr"
)

type Ocr struct {
	APIKey       string // API密钥
	FromDevice   string // 来源
	ClientIP     string // 客户端出口IP
	DetectType   string // OCR接口类型
	LanguageType string // 文字类型
	ImageType    int    // 图片资源类型
	image        string // 图片资源, 文件路径
	Request      *Request
	params       map[string]string
	words        []string // 文本
}

func (o *Ocr) updateParams() {
	if len(o.FromDevice) == 0 {
		o.FromDevice = FROM_DEVICE_API
	}

	if len(o.ClientIP) == 0 {
		o.ClientIP = CLIENT_IP_DEFAULT
	}

	if len(o.DetectType) == 0 {
		o.DetectType = DETECT_TYPE_LOCATE_RECOGNIZE
	}

	if len(o.LanguageType) == 0 {
		o.LanguageType = LANGUAGE_TYPE_CHN_ENG
	}

	if (o.ImageType != IMAGE_TYPE_BASE64) || (o.ImageType != IMAGE_TYPE_ORIGINAL) {
		o.ImageType = IMAGE_TYPE_BASE64
	}

	o.params["fromdevice"] = o.FromDevice
	o.params["clientip"] = o.ClientIP
	o.params["detecttype"] = o.DetectType
	o.params["languagetype"] = o.LanguageType
	o.params["imagetype"] = strconv.Itoa(o.ImageType)
}

// 扫描图片
func (o *Ocr) Scan(filename string) ([]string, error) {
	_, err := o.detectImage(filename)
	if err != nil {
		return nil, err
	}

	o.image = filename

	record, err := o.recognize()
	if err != nil {
		return nil, err
	}

	if record.ErrNum.String() != "0" {
		return nil, errors.New(record.ErrMsg)
	}

	o.words = []string{}
	for _, retData := range record.RetData {
		o.words = append(o.words, retData.Word)
	}

	return o.words, nil
}

// 识别图片
func (o *Ocr) recognize() (*record, error) {
	o.params = map[string]string{}

	if len(o.ClientIP) == 0 {
		o.ClientIP = CLIENT_IP_DEFAULT
	}

	o.updateParams()

	if o.ImageType == IMAGE_TYPE_BASE64 {
		imageStr, err := o.encodeImage()
		if err != nil {
			return nil, err
		}
		o.params[IMAGE_FIELD_NAME] = imageStr
		o.createRequest()
	} else {
		err := o.createFormRequest()
		if err != nil {
			return nil, err
		}
	}

	return o.sendRequest()
}

func (o *Ocr) createRequest() {
	o.Request = NewRequest()
	o.Request.URL = API_URL
	o.Request.SetHeader("apikey", o.APIKey)
	o.Request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	o.Request.WriteBody(o.params)
}

func (o *Ocr) createFormRequest() error {
	o.Request = NewFormRequest()
	o.Request.URL = API_URL
	o.Request.SetHeader("apikey", o.APIKey)
	return o.Request.WriteFormFileBody(o.params, IMAGE_FIELD_NAME, o.image)
}

func (o *Ocr) sendRequest() (*record, error) {
	resp, err := o.Request.Post()

	if err != nil {
		return nil, err
	}

	return parseResponse(resp)
}

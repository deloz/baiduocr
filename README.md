# baiduocr

百度OCR文字识别[API](http://apistore.baidu.com/apiworks/servicedetail/146.html) For Go

## Installation

```go
go get -u github.com/deloz/baiduocr
```

## 参数说明

``` go
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
````

## 使用例子

```go

package main

import (
	"log"
	"os"

	"github.com/deloz/baiduocr"
)

func main() {
	ocr := &baiduocr.Ocr{
		APIKey:       "your app key",
		
		//如果以下参数省略,会自动使用默认值
		FromDevice:   baiduocr.FROM_DEVICE_API,
		ClientIP:     "your client IP",
		DetectType:   baiduocr.DETECT_TYPE_LOCATE_RECOGNIZE,
		LanguageType: baiduocr.LANGUAGE_TYPE_CHN_ENG,
		ImageType:    baiduocr.IMAGE_TYPE_ORIGINAL,
		//或者 使用base64编码
		//ImageType: baiduocr.IMAGE_TYPE_BASE64,
	}

	words, err := ocr.Scan("x12306.jpg")
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
```

## 结果对比

>1. 测试结果 (截于京东注册页面)：

![测试图片](https://github.com/deloz/baiduocr/blob/master/examples/jd.jpg?raw=true) 

```
）手机快速注册  中国大陆手机用户，编辑短信“jD”发送到：1069013355500:
```

>2 测试结果 (截于12306, 图片比较模糊)：

![测试图片](https://github.com/deloz/baiduocr/blob/master/examples/12306.jpg?raw=true) 

```
欢
项服务
1服务条款确认
您点击服务条款页面下的我同意按钮，EmU您已阅读、了解并完全同意服务条款中的各项内容，包舌本网站对服务条款所作的任何修改，除
另
2服务条款的房改：
本网站在必要时可修改服务条款，并在网站进行公告，一经公告，立胜效，如您继续使服务，贝把y您已接受修订的服务条款。
3
考虑至本网站胪服务的重要性，您同意在注册时提供真实‘完整及准确的个人资料，并及时更新。
如您提供的资料不准确，或本团时占有合理的理由认为该资料不真实，不完整｀不准确，本团时占有权暂停或终止您的注册身份及资料，并拒绝您使
用本网站的服务。
4月护资料及保密：
注册时，请您选择填写胪名和它码’并按页面提示提交相关信息，您负有对J胪名和它码保密的义务’并对1刻胪名和它码下发生的所有活动
  承担责任，您同意邮件服务和手机9信的服务的使舳您自己手担破，本网站不会向您所使用服务所涉及相关方之外的其他方公开或透露您的个人
资料，除法律规定。
5责任自晚斜明t制
  （D遇以下情况，本网站不学担任何责任，包括但不仅限于：
①因不可抗力、系统故障、通讯故障、网各拥堵｀供电系统故障、恶意攻击等造成本网站未能及时‘准确完整地提供服务，
②无论在任何原因下，您通过使肘网占上的信息或由本团时占链接的其他网占上的信息，或其他与本团时占链接的网占上的信息所导致的任何损失
  或损害
G在·3J胪注册第二款情形下，注册胪被暂停使用以及因此导致已勾车票不能在本网站改签退票等后果。
  （2）本网站负责对本网站上的信息进行审核与更新，但并不就信息白时效性准确性以及服务功能的完整性和可靠性承担任何义务1呢偿责
```


## Test

```go
go test
```

## Contributing

1. Fork it ( https://github.com/deloz/baiduocr/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## LICENSE

The MIT License (MIT) Copyright (c) 2015 Deloz


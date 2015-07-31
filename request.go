package baiduocr

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Request struct {
	URL     string
	Params  map[string]string // 请求参数
	Body    io.Reader         // 请求实体
	Headers map[string]string
}

func (r *Request) Post() (*http.Response, error) {
	req, err := http.NewRequest("POST", r.URL, r.Body)
	if err != nil {
		return nil, err
	}

	for key, val := range r.Headers {
		req.Header.Add(key, val)
	}

	client := &http.Client{}

	return client.Do(req)
}
func (r *Request) SetHeader(name, value string) {
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}

	r.Headers[name] = value
}

func (r *Request) SetParam(name, value string) {
	if r.Params == nil {
		r.Params = map[string]string{}
	}

	r.Params[name] = value
}

func (r *Request) SetBody(body io.Reader) {
	r.Body = body
}

func (r *Request) WriteBody(params map[string]string) {
	v := url.Values{}
	for name, value := range params {
		v.Set(name, value)
	}
	r.Body = strings.NewReader(v.Encode())
}

func (r *Request) WriteFormFileBody(params map[string]string, fieldName string, filename string) error {
	body, contentType, err := CreateFormFileBody(params, fieldName, filename)
	if err != nil {
		return err
	}

	r.Body = body

	r.SetHeader("Content-Type", contentType)

	return nil
}

func NewRequest() *Request {
	return &Request{
		Params: map[string]string{},
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:38.0) Gecko/20100101 Firefox/38.0 Iceweasel/38.1.0",
		},
	}
}

func NewFormRequest() *Request {
	req := NewRequest()
	return req
}

func CreateFormFileBody(params map[string]string, fieldName string, filename string) (io.Reader, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, "", err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, "", err
	}

	err = file.Close()
	if err != nil {
		return nil, "", err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, fi.Name())
	if err != nil {
		return nil, "", err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, "", err
	}

	for key, val := range params {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, "", err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}

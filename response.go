package baiduocr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RetData struct {
	Rect struct {
		Left   string `json:left`
		Top    string `json:top`
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"rect"`
	Word string `json:"word"`
}

type record struct {
	ErrNum    json.Number `json:"errNum"`
	ErrMsg    string      `json:"errMsg"`
	QuerySign string      `json:"querySign"`
	RetData   []RetData   `json:"retData"`
}

func parseResponse(resp *http.Response) (*record, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r record
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, err
}

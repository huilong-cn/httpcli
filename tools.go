package httpcli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ziipin-server/niuhe"
)

var EmptyBody []byte = []byte{}
var NonHeader http.Header = Header()

func Header() http.Header {
	return make(http.Header)
}

func Values() url.Values {
	return make(url.Values)
}

const (
	TokenBearer = "Bearer"
	TokenBasic  = "Basic"
)

// BearerHeader return header.Add("Authorization", "Bearer "+token)
// tokenType a
func HeaderWithAuth(tokenType, token string) http.Header {
	autheader := Header()
	autheader.Set(Auth(tokenType, token))
	return autheader
}
func Auth(tokenType, token string) (string, string) {
	return "Authorization", tokenType + " " + token
}

// genGetRequest 生成 Get url raw query request
func GenUrlRequest(url string, values url.Values, extendHeader http.Header) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		niuhe.LogError("genGetRequest url : %s, err: %s", url, err.Error())
		return nil, err
	}
	if values != nil {
		request.URL.RawQuery = values.Encode()
	}
	fillExtendHeader(request, extendHeader)

	return request, nil
}

// genGetRequest 生成 Get url raw query request
func GenFormRequest(url string, values url.Values, extendHeader http.Header) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(values.Encode())))
	if err != nil {
		niuhe.LogError("genGetRequest url : %s, err: %s", url, err.Error())
		return nil, err
	}
	request.Form = values
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	fillExtendHeader(request, extendHeader)

	return request, nil
}

// genJsonRequest  gen application/json post request
func GenJsonRequest(url string, requestData interface{}, extendHeader http.Header) (*http.Request, error) {
	requestBytes, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBytes))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	fillExtendHeader(request, extendHeader)
	return request, nil
}

// fillExtendHeader 填充扩展HEADER
func fillExtendHeader(request *http.Request, extendHeader http.Header) {
	for key, values := range extendHeader {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
}

func DecodeWrap(body []byte, rsp interface{}) error {
	wrapRsp := new(Wrap)
	wrapRsp.Data = rsp

	if err := json.Unmarshal(body, wrapRsp); err != nil {
		niuhe.LogError("DecodeWrap unmarshal  resp.BodyL: '%s' fail: %s", string(body), err.Error())
		return err
	}
	if wrapRsp.Result != 0 {
		if wrapRsp.Result != -1 {
			return niuhe.NewCommError(wrapRsp.Result, wrapRsp.Message)
		}
		return errors.New(wrapRsp.Message)
	}
	return nil
}

func DecodeUnwrap(body []byte, rsp interface{}) error {
	if err := json.Unmarshal(body, rsp); err != nil {
		niuhe.LogError("DecodeUnwrap unmarshal resp.Body: '%s' err: %s", string(body), err.Error())
		return err
	}

	return nil
}

func ReadBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		niuhe.LogError("DecodeWrap status :%s read resp.Body fail: %s", resp.Status, err.Error())
		return EmptyBody, err
	}
	return body, nil
}

func IsStatusOK(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad HTTP Status: %s", resp.Status)
	}
	return nil
}

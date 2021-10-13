package httpcli

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ziipin-server/niuhe"
)

func Get(url string, values url.Values, extendHeader http.Header) ([]byte, error) {
	req, err := GenUrlRequest(url, values, extendHeader)
	if err != nil {
		return EmptyBody, err
	}
	resp, err := do3(req)
	if err != nil {
		return EmptyBody, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		niuhe.LogError("DecodeWrap status :%s read resp.Body fail: %s", resp.Status, err.Error())
		return EmptyBody, err
	}
	return body, err
}

func JsonGet(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	req, err := GenUrlRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := do3(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = IsStatusOK(resp)
	if err != nil {
		return err
	}
	body, err := ReadBody(resp)
	if err != nil {
		return err
	}
	return DecodeWrap(body, rsp)
}

func JsonGetUnwrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	req, err := GenUrlRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := do3(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = IsStatusOK(resp)
	if err != nil {
		return err
	}
	body, err := ReadBody(resp)
	if err != nil {
		return err
	}
	return DecodeUnwrap(body, rsp)
}

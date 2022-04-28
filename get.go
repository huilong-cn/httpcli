package httpcli

import (
	"net/http"
	"net/url"
)

func FormGet(url string, values url.Values, extendHeader http.Header) ([]byte, error) {
	return DefaultHttpClient().FormGet(url, values, extendHeader)
}

func JsonGet(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	return DefaultHttpClient().JsonGet(url, values, rsp, extendHeader)
}

func JsonGetUnwrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	return DefaultHttpClient().JsonGetUnwrap(url, values, rsp, extendHeader)
}

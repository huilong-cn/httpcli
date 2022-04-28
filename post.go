package httpcli

import (
	"net/http"
	"net/url"
)

//Post form http post
func FormPost(url string, values url.Values, extendHeader http.Header) ([]byte, error) {
	return DefaultHttpClient().FormPost(url, values, extendHeader)
}

//Post form http post req[json] wrap(rsp)[json]
func PostWrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	return DefaultHttpClient().PostWrap(url, values, rsp, extendHeader)
}

//PostUnwrap form http post req[json] (rsp)[json]
func PostUnwrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	return DefaultHttpClient().PostUnwrap(url, values, rsp, extendHeader)
}

//JsonPost http post req[json] wrap(rsp)[json]
func JsonPost(url string, req interface{}, rsp interface{}, extendHeader http.Header) error {
	return DefaultHttpClient().JsonPost(url, req, rsp, extendHeader)
}

// JsonPostUnwrap http post json rsp => json
func JsonPostUnwrap(url string, req interface{}, rsp interface{}, extendHeader http.Header) error {
	return DefaultHttpClient().JsonPostUnwrap(url, req, rsp, extendHeader)
}

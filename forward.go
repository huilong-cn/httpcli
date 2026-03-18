package httpcli

import (
	"net/http"
)

// JsonForward http forward json []byte  => rsp []byte
func JsonForward(url string, requestBody []byte, extendHeader http.Header) ([]byte, error) {
	return DefaultHttpClient().JsonForward(url, requestBody, extendHeader)
}

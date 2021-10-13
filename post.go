package httpcli

import (
	"net/http"
	"net/url"

	"github.com/ziipin-server/niuhe"
)

//Post form http post
func Post(url string, values url.Values, extendHeader http.Header) ([]byte, error) {
	request, err := GenFormRequest(url, values, extendHeader)
	if err != nil {
		return EmptyBody, err
	}
	resp, err := do3(request)
	if err != nil {
		return EmptyBody, err
	}
	defer resp.Body.Close()
	err = IsStatusOK(resp)
	if err != nil {
		return EmptyBody, err
	}
	body, err := ReadBody(resp)
	if err != nil {
		niuhe.LogError("PostRpc url : %s req : %v, err : %s", url, values, err)
		return EmptyBody, err
	}
	return body, nil
}

//Post form http post req[json] wrap(rsp)[json]
func PostWrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	request, err := GenFormRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := do3(request)
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
	err = DecodeWrap(body, rsp)
	if err != nil {
		niuhe.LogError("PostRpc url : %s req : %v, err : %s", url, values, err)
		return err
	}
	return nil
}

//PostUnwrap form http post req[json] (rsp)[json]
func PostUnwrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	request, err := GenFormRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := do3(request)
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
	err = DecodeUnwrap(body, rsp)
	if err != nil {
		niuhe.LogError("PostUnwrap url : %s req : %v, err : %s", url, values.Encode(), err)
		return err
	}
	return nil
}

//JsonPost http post req[json] wrap(rsp)[json]
func JsonPost(url string, req interface{}, rsp interface{}, extendHeader http.Header) error {
	request, err := GenJsonRequest(url, req, extendHeader)
	if err != nil {
		return err
	}
	resp, err := do3(request)
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
	err = DecodeWrap(body, rsp)
	if err != nil {
		niuhe.LogError("JsonPost url : %s req : %v, err : %s", url, req, err)
		return err
	}
	return nil
}

// JsonPostUnwrap http post json rsp => json
func JsonPostUnwrap(url string, req interface{}, rsp interface{}, extendHeader http.Header) error {
	request, err := GenJsonRequest(url, req, extendHeader)
	if err != nil {
		return err
	}
	resp, err := do3(request)
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
	err = DecodeUnwrap(body, rsp)
	if err != nil {
		niuhe.LogError("JsonPostUnwrap url : %s req : %v, err : %s", url, req, err)
		return err
	}
	return nil
}

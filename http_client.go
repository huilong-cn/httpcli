package httpcli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/ziipin-server/niuhe"
)

type HttpCli struct {
	*http.Client
}

//Post form http post
func (httpcli *HttpCli) Post(url string, values url.Values, extendHeader http.Header) ([]byte, error) {
	request, err := GenFormRequest(url, values, extendHeader)
	if err != nil {
		return EmptyBody, err
	}
	resp, err := httpcli.do3(request)
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
func (httpcli *HttpCli) PostWrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	request, err := GenFormRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := httpcli.do3(request)
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
func (httpcli *HttpCli) PostUnwrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	request, err := GenFormRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := httpcli.do3(request)
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
func (httpcli *HttpCli) JsonPost(url string, req interface{}, rsp interface{}, extendHeader http.Header) error {
	request, err := GenJsonRequest(url, req, extendHeader)
	if err != nil {
		return err
	}
	resp, err := httpcli.do3(request)
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
func (httpcli *HttpCli) JsonPostUnwrap(url string, req interface{}, rsp interface{}, extendHeader http.Header) error {
	request, err := GenJsonRequest(url, req, extendHeader)
	if err != nil {
		return err
	}
	resp, err := httpcli.do3(request)
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

func (httpcli *HttpCli) FormGet(url string, values url.Values, extendHeader http.Header) ([]byte, error) {
	req, err := GenUrlRequest(url, values, extendHeader)
	if err != nil {
		return EmptyBody, err
	}
	resp, err := httpcli.do3(req)
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

func (httpcli *HttpCli) JsonGet(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	req, err := GenUrlRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := httpcli.do3(req)
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

func (httpcli *HttpCli) JsonGetUnwrap(url string, values url.Values, rsp interface{}, extendHeader http.Header) error {
	req, err := GenUrlRequest(url, values, extendHeader)
	if err != nil {
		return err
	}
	resp, err := httpcli.do3(req)
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

func (httpcli *HttpCli) do3(req *http.Request) (*http.Response, error) {
	return httpcli.DoN(req, MAX_RETRY_TIMES)
}

//DoN retry N when network error, except timeout error
func (httpcli *HttpCli) DoN(req *http.Request, retries int) (*http.Response, error) {
	if retries < 1 { //修正最少次数为1
		retries = 1
	}
	for i := 0; i < retries; i++ {
		rsp, err := httpcli.Do(req)
		if err != nil {
			niuhe.LogError("do request url(%s) error : %s", req.URL, err)
			niuhe.LogError("do retry times:%d", i)
			if urlerr, ok := err.(*url.Error); ok && urlerr.Timeout() { // stop retry when timeout error
				httprequest, _ := httputil.DumpRequest(req, false)
				niuhe.LogError("request url(%s), req(%s) error : %s", req.URL.String(), string(httprequest), err.Error())
				if timeoutCallback != nil {
					timeoutCallback("HTTPCLI-TIMEOUT", fmt.Sprintf("request url(%s), req(%s) error : %s", req.URL.String(), string(httprequest), err.Error()))
				}
				return nil, err
			}
			time.Sleep((time.Second * time.Duration(i+1)))
			continue
		}
		return rsp, nil
	}
	httprequest, _ := httputil.DumpRequest(req, false)
	niuhe.LogError("do retry %d times url:%s failed request:%s", retries, req.URL, string(httprequest))
	if timeoutCallback != nil {
		timeoutCallback("HTTPCLI-RETRY", fmt.Sprintf("request url(%s), req(%s) error max retry", req.URL.String(), string(httprequest)))
	}
	return nil, fmt.Errorf("httpcli DoN:%d request err", retries)
}

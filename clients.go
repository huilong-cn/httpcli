package httpcli

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	httpClient *http.Client
)

type clientFactory struct {
	sync.Mutex
	clients map[string]*HttpCli
}

var defualtClientFactory *clientFactory

// opts[0] proxy
func (cf *clientFactory) get(clientid string, tlsCfg *tls.Config, opts ...string) *HttpCli {
	cf.Lock()
	defer cf.Unlock()
	if client, ok := cf.clients[clientid]; ok {
		return client
	}
	var proxyFunc = http.ProxyFromEnvironment
	if len(opts) > 0 && opts[0] != "" {
		proxyFunc = func(request *http.Request) (*url.URL, error) {
			return url.Parse(opts[0])
		}
	}
	newClient := &HttpCli{
		&http.Client{
			Transport: &http.Transport{
				Proxy: proxyFunc,
				Dial: (&net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 300 * time.Second,
				}).Dial,
				IdleConnTimeout:       300 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConns:          4096,
				MaxConnsPerHost:       1024,
				MaxIdleConnsPerHost:   1024,
			},
			Timeout: time.Second * 30,
		}}
	cf.clients[clientid] = newClient
	return newClient
}

const (
	DefaultHttpClientID = "default_http_client_id"
)

func HttpClient(clientid string, proxystr ...string) *HttpCli {
	return defualtClientFactory.get(clientid, nil, proxystr...)
}

func HttpClientUseTLSConf(clientid string, tlsCfg *tls.Config, proxystr ...string) *HttpCli {
	return defualtClientFactory.get(clientid, tlsCfg, proxystr...)
}

func DefaultHttpClient() *HttpCli {
	return HttpClient(DefaultHttpClientID)
}

// init HTTPClient
func init() {
	defualtClientFactory = &clientFactory{
		clients: make(map[string]*HttpCli),
	}
}

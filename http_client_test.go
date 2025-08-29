package httpcli

import (
	"net/url"
	"testing"
)

func TestJsonGetUnwrap(t *testing.T) {
	resp := make(map[string]interface{})
	err := JsonGetUnwrap("https://ipinfo.io/json", url.Values{}, &resp, NonHeader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestPostUnwrap(t *testing.T) {
	resp := make(map[string]interface{})
	err := PostUnwrap("https://ipinfo.io/json", url.Values{}, &resp, NonHeader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestFormPost(t *testing.T) {
	bytes, err := FormPost("https://google.com", url.Values{}, NonHeader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes)
}

func TestFormGet(t *testing.T) {
	bytes, err := FormGet("https://ipinfo.io/ip", url.Values{}, NonHeader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

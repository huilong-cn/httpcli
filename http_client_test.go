package httpcli

import (
	"net/url"
	"os"
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

func TestPostByte(t *testing.T) {
	bytes, err := PostBytes("https://ipinfo.io/json", []byte{}, NonHeader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestPostByte2(t *testing.T) {
	bytes, err := os.ReadFile("/Users/long/Downloads/turkey_video2.mp4")
	if err != nil {
		t.Fatal(err)
	}
	respBody, err := PostBytes("http://192.168.20.32:8800/api/splitvideo?format=mp3", bytes, NonHeader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(respBody))
}

func TestHead(t *testing.T) {
	err := DefaultHttpClient().HEAD("https://download.sawagames.com/baloot_26-07-02.apk", NonHeader)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(string(bytes))
}

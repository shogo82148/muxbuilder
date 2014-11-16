package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	handler, err := NewHandler()
	if err != nil {
		panic(err)
	}
	ts := httptest.NewServer(handler)

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("GET / failed: %v", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("GET / failed: invalid status code %d", res.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	expected := "Your request method is GET"
	if string(responseBody) != expected {
		t.Errorf("GET / return %s\nwant %s", string(responseBody), expected)
	}
}

func TestPost(t *testing.T) {
	handler, err := NewHandler()
	if err != nil {
		panic(err)
	}
	ts := httptest.NewServer(handler)

	res, err := http.PostForm(ts.URL, url.Values{})
	if err != nil {
		t.Errorf("POST / failed: %v", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("POST / failed: invalid status code %d", res.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	expected := "Your request method is POST"
	if string(responseBody) != expected {
		t.Errorf("POST / return %s\nwant %s", string(responseBody), expected)
	}
}

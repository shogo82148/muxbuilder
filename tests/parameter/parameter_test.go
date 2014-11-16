package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	URL      string
	Response string
}

var testCases = []TestCase{
	TestCase{
		URL:      "/foo/hogehoge/bar",
		Response: "My name is hogehoge",
	},
}

func TestHandler(t *testing.T) {
	handler, err := NewHandler()
	if err != nil {
		panic(err)
	}
	ts := httptest.NewServer(handler)

	for _, testCase := range testCases {
		url := ts.URL + testCase.URL
		res, err := http.Get(url)
		if err != nil {
			t.Errorf("GET %s failed: %v", testCase.URL, err)
			continue
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("GET %s failed: invalid status code %d", testCase.URL, res.StatusCode)
		}

		responseBody, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if string(responseBody) != testCase.Response {
			t.Errorf("GET %s return %s\nwant %s", testCase.URL, string(responseBody), testCase.Response)
		}
	}
}

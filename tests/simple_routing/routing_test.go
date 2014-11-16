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
		URL:      "/",
		Response: "This is Root handler",
	},
	TestCase{
		URL:      "/foo",
		Response: "This is Foo handler",
	},
	TestCase{
		URL:      "/bar",
		Response: "This is Bar handler",
	},
}

func TestHandler(t *testing.T) {
	handler, err := NewHandler()
	if err != nil {
		panic(err)
	}
	ts := httptest.NewServer(handler)

	for _, testCase := range testCases {
		res, err := http.Get(ts.URL + testCase.URL)
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

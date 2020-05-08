// +build unit

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var listAddResult = []AddResult{
	{1, 1, 2},
}

func TestAdd(t *testing.T) {
	for _, test := range listAddResult {
		result := add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")

	if err != nil {
		t.Fatal("Couldn't open file")
	}

	if string(data) != "hello world" {
		t.Fatal("String contents do not match expected")
	}
}

func TestHTPPRequest(t *testing.T) {
	handler := func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "{\"status\" : \"good\"}")
	}

	req := httptest.NewRequest("GET", "http://google.com", nil)
	res := httptest.NewRecorder()
	handler(res, req)

	resp := res.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		t.Fatal("Status code Not OK")
	}

}

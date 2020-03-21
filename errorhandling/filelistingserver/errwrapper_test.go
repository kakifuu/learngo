package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserErr string

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noErr, 200, "No Error"},
}

func (e testingUserErr) Error() string {
	return e.Message()
}

func (e testingUserErr) Message() string {
	return string(e)
}

func errPanic(_ http.ResponseWriter, _ *http.Request) error {
	panic(123)
}

func errUserError(_ http.ResponseWriter, _ *http.Request) error {
	return testingUserErr("user error")
}

func errNotFound(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter, _ *http.Request) error {
	return errors.New("unknown error")
}

func noErr(writer http.ResponseWriter, _ *http.Request) error {
	fmt.Fprintln(writer, "No Error")
	return nil
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("Expect (%d, %s), got (%d, %s)",
			expectedCode, expectedMsg, resp.StatusCode, body)
	}
}

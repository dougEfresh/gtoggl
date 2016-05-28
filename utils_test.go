package gtoggl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type TestLogger struct {
}

var _, debugMode = os.LookupEnv("GTOGGL_TEST_DEBUG")

func (l *TestLogger) Printf(format string, v ...interface{}) {
	if debugMode {
		fmt.Printf(format, v)
	}
}

var mockResponses = make(map[string][]byte)

var testLogger = &TestLogger{}

type mockFunc func(req *http.Request) []byte

type mockTransport struct {
	mock mockFunc
}

func newMockTransport(f mockFunc) http.RoundTripper {
	return &mockTransport{mock: f}
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		mockResponses[path], err = ioutil.ReadFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func getResponse() mockFunc {
	return func(req *http.Request) []byte {
		r := fmt.Sprintf("mock/%s%s.json", req.Method, req.URL.Path)
		resp := mockResponses[r]
		if resp == nil {
			panic("Unknown request " + r)
		}
		return resp
	}
}

func load() {
	if len(mockResponses) > 0 {
		return
	}
	err := filepath.Walk("mock", visit)
	if err != nil {
		panic(err)
	}
}

var nullResponse = []byte("")

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create mocked http.Response
	response := &http.Response{Header: make(http.Header), Request: req, StatusCode: http.StatusOK}
	response.Header.Set("Content-Type", "application/json")
	response.Header.Set("Set-Cookie", "toggl_api_session_new=MTM2MzA4MJa8jA3OHxEdi1CQkFFQ180SUFBUkFCRUFBQVlQLUNBQUVHYzNSeWFXNW5EQXdBQ25ObGMzTnBiMjVmYVdRR2MzUnlhVzVuREQ0QVBIUnZaMmRzTFdGd2FTMXpaWE56YVc5dUxUSXRaalU1WmpaalpEUTVOV1ZsTVRoaE1UaGhaalpqWkRkbU5XWTJNV0psWVRnd09EWmlPVEV3WkE9PXweAkG7kI6NBG-iqvhNn1MSDhkz2Pz_UYTzdBvZjCaA==; Path=/; Expires=Wed, 13 Mar 2013 09:54:38 UTC; Max-Age=86400; HttpOnly")
	if strings.Contains(req.URL.Path, "/sessions") {
		response.Body = ioutil.NopCloser(bytes.NewReader(nullResponse))
		return response, nil
	}
	responseBody := t.mock(req)
	response.Body = ioutil.NopCloser(bytes.NewReader(responseBody))
	return response, nil
}

func mockClient() *TogglHttpClient {
	load()
	httpClient := &http.Client{Transport: newMockTransport(getResponse())}
	optionsWithClient := []ClientOptionFunc{SetHttpClient(httpClient)}
	client, err := NewClient("abc1234567890def", optionsWithClient...)
	if err != nil {
		panic(err)
	}
	return client
}

func mockClientOptions(m mockFunc, options []ClientOptionFunc) *TogglHttpClient {
	httpClient := &http.Client{Transport: newMockTransport(m)}
	optionsWithClient := options[0 : len(options)+1]
	optionsWithClient[len(options)] = SetHttpClient(httpClient)
	client, err := NewClient("abc1234567890def", optionsWithClient...)
	if err != nil {
		panic(err)
	}
	return client
}

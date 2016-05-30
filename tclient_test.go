package gtoggl

import (
	"bytes"
	"gopkg.in/dougEfresh/toggl-http-client.v8"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockTransport struct {
}

var nullResponse = []byte("")

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create mocked http.Response
	response := &http.Response{Header: make(http.Header), Request: req, StatusCode: http.StatusOK}
	response.Header.Set("Content-Type", "application/json")
	response.Header.Set("Set-Cookie", "toggl_api_session_new=MTM2MzA4MJa8jA3OHxEdi1CQkFFQ180SUFBUkFCRUFBQVlQLUNBQUVHYzNSeWFXNW5EQXdBQ25ObGMzTnBiMjVmYVdRR2MzUnlhVzVuREQ0QVBIUnZaMmRzTFdGd2FTMXpaWE56YVc5dUxUSXRaalU1WmpaalpEUTVOV1ZsTVRoaE1UaGhaalpqWkRkbU5XWTJNV0psWVRnd09EWmlPVEV3WkE9PXweAkG7kI6NBG-iqvhNn1MSDhkz2Pz_UYTzdBvZjCaA==; Path=/; Expires=Wed, 13 Mar 2013 09:54:38 UTC; Max-Age=86400; HttpOnly")
	response.Body = ioutil.NopCloser(bytes.NewReader(nullResponse))
	return response, nil
}

func mockClient(t *testing.T) *TogglClient {
	httpClient := &http.Client{Transport: &mockTransport{}}
	client, err := NewClient("abc1234567890def", ghttp.SetHttpClient(httpClient))
	if err != nil {
		panic(err)
	}
	return client
}

func togglClient(t *testing.T) *TogglClient {
	client := mockClient(t)
	return client
}

func TestClientCreate(t *testing.T) {
	tClient := togglClient(t)
	if tClient.UserClient == nil {
		t.Fatal("No User Client")
	}
	if tClient.TClient == nil {
		t.Fatal("No TClient")
	}
	if tClient.WorkspaceClient == nil {
		t.Fatal("No Workspace")
	}
	if tClient.ProjectClient == nil {
		t.Fatal("No Pc")
	}
	if tClient.TimeentryClient == nil {
		t.Fatal("No Te")
	}
}

package gtoggl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DefaultAuthPassword = "api_token"
	DefaultMaxRetries   = 5
	DefaultGzipEnabled  = false
	DefaultUrl          = "https://www.toggl.com/api/v8"
	DefaultVersion      = "v8"
	TogglCreator        = "github.com/dougEresh/gtoggl"
)

// Client is an Toggl REST client. Created by calling NewClient.
type TogglHttpClient struct {
	client        *http.Client // net/http Client to use for requests
	version       string       // v8
	Url           string       // set of URLs passed initially to the client
	errorLog      Logger       // error log for critical messages
	infoLog       Logger       // information log for e.g. response times
	traceLog      Logger       // trace log for debugging
	password      string       // password for HTTP Basic Auth
	maxRetries    uint
	sessionCookie string //24 hour session cookie
	gzipEnabled   bool   // gzip compression enabled or disabled (default)
}

type TogglError struct {
	Code   int
	Status string
	Msg    string
}

type TogglResponse struct {
	Data *json.RawMessage `json:"data"`
}

func (e *TogglError) Error() string {
	return fmt.Sprintf("%s\t%s\n", e.Status, e.Msg)
}

// ClientOptionFunc is a function that configures a Client.
// It is used in NewClient.
type ClientOptionFunc func(*TogglHttpClient) error

// Return a new TogglHttpClient . An error is also returned when some configuration option is invalid
//    tc,err := gtoggl.NewClient("token")
func NewClient(key string, options ...ClientOptionFunc) (*TogglHttpClient, error) {
	// Set up the client
	c := &TogglHttpClient{
		client:      http.DefaultClient,
		maxRetries:  DefaultMaxRetries,
		Url:         DefaultUrl,
		version:     DefaultVersion,
		gzipEnabled: DefaultGzipEnabled,
		password:    DefaultAuthPassword,
	}

	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	if len(key) < 1 {
		return nil, errors.New("Token required")
	}

	_, err := c.authenticate(key)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// SetHttpClient can be used to specify the http.Client to use when making
// HTTP requests to Toggl
func SetHttpClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *TogglHttpClient) error {
		if httpClient != nil {
			c.client = httpClient
		} else {
			c.client = http.DefaultClient
		}
		return nil
	}
}

// SetURL defines the base URL. See DefaultUrl
func SetURL(url string) ClientOptionFunc {
	return func(c *TogglHttpClient) error {
		switch len(url) {
		case 0:
			c.Url = DefaultUrl
		default:
			c.Url = url
		}
		return nil
	}
}

//Custom logger to print HTTP requests
func SetTraceLogger(l Logger) ClientOptionFunc {
	return func(c *TogglHttpClient) error {
		c.traceLog = l
		return nil
	}
}

//Custom logger to handle error messages
func SetErrorLogger(l Logger) ClientOptionFunc {
	return func(c *TogglHttpClient) error {
		c.errorLog = l
		return nil
	}
}

//Custom logger to handle info messages
func SetInfoLogger(l Logger) ClientOptionFunc {
	return func(c *TogglHttpClient) error {
		c.infoLog = l
		return nil
	}
}

// Returns the session cookie
func (c *TogglHttpClient) String() string {
	return fmt.Sprintf("{sessionCookie=%s}", c.sessionCookie)
}

func (c *TogglHttpClient) authenticate(key string) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.Url, "sessions"), nil)
	if err != nil {
		return nil, err
	}
	c.dumpRequest(req)
	req.SetBasicAuth(key, "api_token")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	c.dumpResponse(resp)
	cookies := resp.Cookies()
	for _, value := range cookies {
		if value.Name == "toggl_api_session_new" {
			c.sessionCookie = value.Value
		}
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, &TogglError{Code: resp.StatusCode, Status: resp.Status, Msg: string(b)}
	}

	return nil, nil
}

var cookieJar = make(map[string]*http.Cookie, 10)

func request(c *TogglHttpClient, method, endpoint string, body []byte) (*json.RawMessage, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	cookie := cookieJar[c.sessionCookie]
	if cookie == nil {
		cookie = &http.Cookie{}
		cookie.Name = "toggl_api_session_new"
		cookie.Value = c.sessionCookie
		cookieJar[c.sessionCookie] = cookie
	}
	req.AddCookie(cookie)
	c.dumpRequest(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	c.dumpResponse(resp)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, &TogglError{Code: resp.StatusCode, Status: resp.Status, Msg: string(b)}
	}
	var raw json.RawMessage
	err = json.Unmarshal(b, &raw)
	if err != nil {
		return nil, err
	}
	return &raw, err
}

// Utility to POST requests
func (c *TogglHttpClient) PostRequest(endpoint string, body []byte) (*json.RawMessage, error) {
	return request(c, "POST", endpoint, body)
}

// Utility to DELETE requests
func (c *TogglHttpClient) DeleteRequest(endpoint string, body []byte) (*json.RawMessage, error) {
	return request(c, "DELETE", endpoint, body)
}

// Utility to PUT requests
func (c *TogglHttpClient) PutRequest(endpoint string, body []byte) (*json.RawMessage, error) {
	return request(c, "PUT", endpoint, body)
}

// Utility to GET requests
func (c *TogglHttpClient) GetRequest(endpoint string) (*json.RawMessage, error) {
	return request(c, "GET", endpoint, nil)
}

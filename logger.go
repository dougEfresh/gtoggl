package gtoggl

import (
	"net/http"
	"net/http/httputil"
)

// Logger specifies the interface for all log operations.
type Logger interface {
	Printf(format string, v ...interface{})
}

// errorf logs to the error log.
func (c *TogglHttpClient) errorf(format string, args ...interface{}) {
	if c.errorLog != nil {
		c.errorLog.Printf(format, args...)
	}
}

// infof logs informational messages.
func (c *TogglHttpClient) infof(format string, args ...interface{}) {
	if c.infoLog != nil {
		c.infoLog.Printf(format, args...)
	}
}

// tracef logs to the trace log.
func (c *TogglHttpClient) tracef(format string, args ...interface{}) {
	if c.traceLog != nil {
		c.traceLog.Printf(format, args...)
	}
}

// dumpRequest dumps the given HTTP request to the trace log.
func (c *TogglHttpClient) dumpRequest(r *http.Request) {
	if c.traceLog != nil {
		out, err := httputil.DumpRequestOut(r, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}

}

// dumpResponse dumps the given HTTP response to the trace log.
func (c *TogglHttpClient) dumpResponse(resp *http.Response) {
	if c.traceLog != nil {
		out, err := httputil.DumpResponse(resp, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}
}

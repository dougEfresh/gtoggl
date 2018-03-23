package gtoggl

import (
	"github.com/dougEfresh/gtoggl-api/gtclient"
	"github.com/dougEfresh/gtoggl-api/gthttp"
	"github.com/dougEfresh/gtoggl-api/gtproject"
	"github.com/dougEfresh/gtoggl-api/gttimentry"
	"github.com/dougEfresh/gtoggl-api/gtuser"
	"github.com/dougEfresh/gtoggl-api/gtworkspace"
)

// Client is an Toggl REST client. Created by calling NewClient.
type TogglClient struct {
	TogglHttpClient *gthttp.TogglHttpClient
	WorkspaceClient *gtworkspace.WorkspaceClient
	ProjectClient   *gtproject.ProjectClient
	TClient         *gtclient.TClient
	UserClient      *gtuser.UserClient
	TimeentryClient *gttimeentry.TimeEntryClient
}

// Return a new TogglHttpClient . An error is also returned when some configuration option is invalid
//    tc,err := gtoggl.NewClient("token")
func NewClient(key string, options ...gthttp.ClientOptionFunc) (*TogglClient, error) {
	// Set up the client
	c, err := gthttp.NewClient(key, options...)
	if err != nil {
		return nil, err
	}
	th := &TogglClient{TogglHttpClient: c,
		WorkspaceClient: gtworkspace.NewClient(c),
		UserClient:      gtuser.NewClient(c),
		ProjectClient:   gtproject.NewClient(c),
		TClient:         gtclient.NewClient(c),
		TimeentryClient: gttimeentry.NewClient(c),
	}
	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return th, nil
}

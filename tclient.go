package gtoggl

import (
	"gopkg.in/dougEfresh/toggl-client.v8"
	"gopkg.in/dougEfresh/toggl-http-client.v8"
	"gopkg.in/dougEfresh/toggl-project.v8"
	"gopkg.in/dougEfresh/toggl-timeentry.v8"
	"gopkg.in/dougEfresh/toggl-user.v8"
	"gopkg.in/dougEfresh/toggl-workspace.v8"
)

// Client is an Toggl REST client. Created by calling NewClient.
type TogglClient struct {
	TogglHttpClient *ghttp.TogglHttpClient
	WorkspaceClient *gworkspace.WorkspaceClient
	ProjectClient   *gproject.ProjectClient
	TClient         *gclient.TClient
	UserClient      *guser.UserClient
	TimeentryClient *gtimeentry.TimeEntryClient
}

// Return a new TogglHttpClient . An error is also returned when some configuration option is invalid
//    tc,err := gtoggl.NewClient("token")
func NewClient(key string, options ...ghttp.ClientOptionFunc) (*TogglClient, error) {
	// Set up the client
	c, err := ghttp.NewClient(key, options...)
	if err != nil {
		return nil, err
	}
	th := &TogglClient{TogglHttpClient: c,
		WorkspaceClient: gworkspace.NewClient(c),
		UserClient:      guser.NewClient(c),
		ProjectClient:   gproject.NewClient(c),
		TClient:         gclient.NewClient(c),
		TimeentryClient: gtimeentry.NewClient(c),
	}
	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return th, nil
}

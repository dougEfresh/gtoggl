// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Command Line Interface to toggl

Example:
	go run gtoggl.go  -t $TOGGL_API -c user
	go run gtoggl.go  -t $TOGGL_API -c user update '{"id":2269184, "email":"newEmail@gmail.com","fullname":"New Name"}'
	go run gtoggl.go  -t $TOGGL_API -d -c project  # debug
*/
package main

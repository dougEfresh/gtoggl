// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gtoggl access to toggl REST API

Example:
        import "gopkg.in/dougEfresh/gtoggl.v8"

        func main() {
       	    tc, err := gtoggl.NewClient("token")
	    if err == nil {
    		panic(err)
            }
            timeentry, _:= tc.TimeentryClient.Get(1)
	}
*/
package gtoggl

// Copyright 2016 Douglas Chimento.  All rights reserved.

/*
Package gtoggl access to toggl REST API

Example:
        import "gopkg.in/dougEfresh/gtoggl.v8"
        import "ggopkg.in/dougEfresh/toggl-timeentry.v8"

        func main() {
	    thc, err := gtoggl.NewClient("token")
	    ...
	    tc, err := gtimeentry.NewClient(thc)
	    ...
	    timeentry,err := tc.Get(1)
	    if err == nil {
		panic(err)
	    }
	}
*/
package gtoggl

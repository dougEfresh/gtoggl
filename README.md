# Gtoggl 
 
 Gtoggl is a [toggl](https://github.com/toggl/toggl_api_docs) client for the [Go](http://www.golang.org/) programming language.
 
[![Build Status](https://travis-ci.org/dougEfresh/gtoggl.svg?branch=master)](https://travis-ci.org/dougEfresh/gtoggl)
[![Go Report Card](https://goreportcard.com/badge/github.com/dougEfresh/gtoggl)](https://goreportcard.com/report/github.com/dougEfresh/gtoggl)
[![GoDoc](https://godoc.org/github.com/dougEfresh/gtoggl?status.svg)](https://godoc.org/github.com/dougEfresh/gtoggl)
[![Coverage Status](https://coveralls.io/repos/github/dougEfresh/gtoggl/badge.svg?branch=master)](https://coveralls.io/github/dougEfresh/gtoggl?branch=master)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dougEfresh/gtoggl/master/LICENSE)

**Example:**

```sh
go get gopkg.in/dougEfresh/gtoggl.v8
```


```go
import "gopkg.in/dougEfresh/gtoggl.v8"

func main() {
    tc, err := gtoggl.NewClient("token")
    if err == nil {
    	panic(err)
    }
    te, _:= tc.TimeentryClient.Get(1)
}
``` 

## Supported Endpoints 

 [Time Entry](https://github.com/dougEfresh/toggl-timeeentry)
 [User](https://github.com/dougEfresh/toggl-user)
 [Workspace](https://github.com/dougEfresh/toggl-workspace)
 [Project](https://github.com/dougEfresh/toggl-project)
 [Client](https://github.com/dougEfresh/toggl-client)
 
 
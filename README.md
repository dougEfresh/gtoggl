# Toggl API for golang

Throttle API for [toggle](https://github.com/toggl/toggl_api_docs/blob/master/toggl_api.md)

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report][report-img]][report]

## Installation 
```shell
$ go get -u github.com/dougEfresh/gtoggl-api
$ go get -u github.com/dougEfresh/gtoggl
```

## Quick Start

```go
import "github.com/dougEfresh/gtoggl"
import "github.com/dougEfresh/gtoggl-api/gtproject"

func main() {
  thc, err := gtoggl.NewClient("token")
  ...
  tc, err := gtproject.NewClient(thc)
  ...
  project,err := tc.Get(1)
  if err == nil {
    panic(err)
   }
}
```


The gtoggl clients provides throttling

## Usage 

See [gtoggl cli](https://github.com/dougEfresh/gtoggl/gtoggl)

## Examples
    
See [godoc][doc] for more examples


## Prerequisites

go 1.x

## Tests
    
```shell
$ go test -v ./...

```


## Deployment

## Contributing
 All PRs are welcome

## Authors

* **Douglas Chimento**  - [dougEfresh][me]

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

### TODO 

[doc-img]: https://godoc.org/github.com/dougEfresh/gtoggl?status.svg
[doc]: https://godoc.org/github.com/dougEfresh/gtoggl
[ci-img]: https://travis-ci.org/dougEfresh/gtoggl.svg?branch=master
[ci]: https://travis-ci.org/dougEfresh/gtoggl
[cov-img]: https://codecov.io/gh/dougEfresh/gtoggl/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/dougEfresh/gtoggl
[glide.lock]: https://github.com/uber-go/zap/blob/master/glide.lock
[zap]: https://github.com/uber-go/zap
[me]: https://github.com/dougEfresh
[report-img]: https://goreportcard.com/badge/github.com/dougEfresh/gtoggl
[report]: https://goreportcard.com/report/github.com/dougEfresh/gtoggl

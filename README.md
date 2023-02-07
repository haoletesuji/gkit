# Gkit

Template for REST/GRPC API server made with Golang.

Features:

- The log uses logrus
- Load middleware cross-domain, add header
- Graceful stop
- Close connection to DB before stop
- Response template error/success/success with pagination

## Installation

To install `gkit` package, you need to install Go.

1. You first need [Go](https://golang.org/) installed then you can use the below Go command to install `gkit`.

```sh
go get -u github.com/haoletesuji/gkit
```

2. Import it in your code:

```go
import "github.com/haoletesuji/gkit"
```

## Quick start

### Starting HTTP server

The example how to use `gkit` with to start REST api in [example](./example/http) folder.

```go
package main

import (
	"example/wire"
	"log"
)

func main() {
	server, err := wire.InitializeServer("gkit_example")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve()
}

```

```
# run main.go and visit 0.0.0.0:3000/api/health on browser
$ go run main.go
```

Pre defination response template:

- Error response: `GET` /api/error

```go
type ErrResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error"`
}
```

```json
{
  "success": false,
  "error": "response error"
}
```

- Successfull response: `GET` /api/success

```go
type SuccessResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
}
```

```json
{
  "success": true,
  "data": [
    {
      "uid": "1",
      "name": "John",
      "profile_image": "https://google.com/image/1"
    }
  ]
}
```

- Successfull with pagination: `GET` /api/success_paging

```go
type Pagination struct {
	Total int64 `json:"total"`
	Limit int64 `json:"limit"`
}

type SuccessPagingResponse[T any] struct {
	Success    bool       `json:"success" example:"true"`
	Data       T          `json:"data"`
	Pagination Pagination `json:"pagination"`
}
```

```json
{
  "success": true,
  "data": [
    {
      "uid": "1",
      "name": "John",
      "profile_image": "https://google.com/image/1"
    }
  ],
  "pagination": {
    "total": 1,
    "limit": 10
  }
}
```

### Starting GRPC server

The example how to use `gkit` with to start grpc server in [example](./example/grpc) folder.
```
# run main.go
$ go run main.go
```

Then you will see the following logging

```
INFO[0000] grpc server listening on   0.0.0.0:50001      type=grpc
```

Using Postman to invoke functions
![Alt text](./images/postman_1.png?raw=true "postman_1")


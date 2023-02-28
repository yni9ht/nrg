# NRG
`nrg`(Ni9ht Router Group) is a simple and easy to use router for your web application. It is written in pure Golang and has no dependencies.

## Features
- [ ] HTTP Methods
  - [x] add http handler
- [ ] Context
  - [x] GetQuery
  - [ ] PostFrom
  - [ ] GetBody
  - [ ] GetHeader
  - [ ] Get
  - [x] JSON
  - [ ] HTML
  - [ ] File
- [ ] Redirect
- [ ] Router
  - [x] Add Router
  - [ ] Static Route
  - [ ] Parameter Route
  - [ ] Regex Route
  - [ ] Group Route
- [ ] Middleware

## Quick Start
1. Installation
```bash
go get -u github.com/yni9ht/nrg
```

2. Example
```go
package main

import (
  "fmt"
  "github.com/yni9ht/nrg"
)

func main() {
  server := nrg.NewServer()

  server.GET("/ping", func(context *nrg.Context) {
    context.JSON(200, "pong")
  })

  if err := server.Run(); err != nil {
    fmt.Printf("error %+v \n", err)
  }
}
```
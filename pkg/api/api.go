package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  api "github.com/woojiahao/govid-19/pkg/api/action"
)

var endpoints = []Endpoint{
  {
    RequestType: GET,
    Path:        "/ping",
    Action:      api.Ping,
  },
  {
    RequestType: GET,
    Path:        "/all",
    Action:      api.All,
  },
}

func Build(engine *gin.Engine) {
  fmt.Println("Building endpoints...")
  for _, endpoint := range endpoints {
    path, action := endpoint.Path, endpoint.Action
    switch endpoint.RequestType {
    case GET:
      engine.GET(path, action)
    case POST:
      engine.POST(path, action)
    case PUT:
      engine.PUT(path, action)
    case DELETE:
      engine.DELETE(path, action)
    }
  }
}

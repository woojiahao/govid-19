package api

import (
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/data"
)

var endpoints = []Endpoint{
  {
    RequestType: GET,
    Path:        "/ping",
    Action: func(c *gin.Context) {
      OK(c, gin.H{"message": "pong"})
    },
  },
  {
    RequestType: GET,
    Path:        "/countries",
    Action: func(c *gin.Context) {
      OK(c, data.Countries)
    },
  },
  {
    RequestType: GET,
    Path:        "/all",
    Action:      All,
  },
}

func Build(engine *gin.Engine) {
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

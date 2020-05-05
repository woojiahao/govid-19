package api

import (
  "github.com/gin-gonic/gin"
)

var endpoints = []Endpoint{
  {
    RequestType: GET,
    Path:        "/ping",
    Action:      Ping,
  },
  {
    RequestType: GET,
    Path:        "/all",
    Action:      All,
  },
  {
    RequestType: GET,
    Path:        "/countries",
    Action:      GetAvailableCountries,
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

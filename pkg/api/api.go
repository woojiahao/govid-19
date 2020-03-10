package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/data"
)

func ping(c *gin.Context) {
  OK(c, gin.H{
    "message": "pong",
  })
}

func all(c *gin.Context) {
  confirmed, deaths, recovered := data.GetAll()

  params := c.Request.URL.Query()
  country, state := params.Get("country"), params.Get("state")
  if country != "" {

  }

  OK(c, gin.H{
    "confirmed": confirmed,
    "deaths":    deaths,
    "recovered": recovered,
  })
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

var endpoints = []Endpoint{
  {
    RequestType: GET,
    Path:        "/ping",
    Action:      ping,
  },
  {
    RequestType: GET,
    Path:        "/all",
    Action:      all,
  },
}

package api

import (
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/database"
)

var databaseManager *database.Manager

var endpoints = []Endpoint{
  {
    GET,
    "/ping",
    func(c *gin.Context) {
      OK(c, gin.H{"message": "pong"})
    },
  },
  {
    GET,
    "/countries",
    getCountries,
  },
  {
    GET,
    "/all",
    All,
  },
  {
    GET,
    "/stats/general",
    getGeneralCountryInformation,
  },
  {
    GET,
    "/stats/overview/:country_id",
    getGeneralCountryInformation,
  },
}

func Build(engine *gin.Engine, manager *database.Manager) {
  databaseManager = manager
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

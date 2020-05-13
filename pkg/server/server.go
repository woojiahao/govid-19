package server

import (
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/api"
  "github.com/woojiahao/govid-19/pkg/data"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "log"
)

func Start() {
  r := gin.Default()

  // CORS configuration must occur before creating any API routes
  log.Print("Setting up CORS")
  r.Use(cors.New(cors.Config{
    AllowAllOrigins: true,
    AllowMethods:    []string{"GET", "OPTIONS"},
    AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
    ExposeHeaders:   []string{"Content-Length"},
  }))

  log.Print("Building API endpoints")
  api.Build(r)

  log.Print("Creating timer to update data daily")
  go Run(data.LoadData)

  log.Print("Running server")
  err := r.Run()
  Check(err)
}

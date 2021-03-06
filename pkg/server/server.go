package server

import (
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/api"
  "github.com/woojiahao/govid-19/pkg/data"
  "github.com/woojiahao/govid-19/pkg/database"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "log"
)

func Start(databaseManager *database.Manager) {
  r := gin.Default()

  // Setting up Gin error handling
  log.Print("Setting up error handling middleware")
  r.Use(handleError())

  // CORS configuration must occur before creating any API routes
  log.Print("Setting up CORS")
  r.Use(cors.New(cors.Config{
    AllowAllOrigins: true,
    AllowMethods:    []string{"GET", "OPTIONS"},
    AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
    ExposeHeaders:   []string{"Content-Length"},
  }))

  log.Print("Building API endpoints")
  api.Build(r, databaseManager)

  log.Print("Creating timer to update data daily")
  go Run(data.Update)

  log.Print("Running server")
  err := r.Run()
  Check(err)
}

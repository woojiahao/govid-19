package server

import (
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/api"
  "github.com/woojiahao/govid-19/pkg/data"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "log"
  "os"
)

// If the current instance already contains the repository, pull the latest changes from the repository
// If the current instance does not contain the repository, clone the repository
func loadData() {
  if _, err := os.Stat(data.Root.AsString()); os.IsNotExist(err) {
    tmp := os.Mkdir(data.Root.AsString(), 0755)
    Check(tmp)

    data.Load()
  } else {
    data.Update()
  }
}

func Start() {
  log.Print("Starting server")
  r := gin.Default()

  log.Print("Loading data")
  loadData()

  log.Print("Building API endpoints")
  api.Build(r)

  r.Use(cors.New(cors.Config{
    AllowAllOrigins: true,
    AllowMethods:    []string{"GET", "PATCH"},
    AllowHeaders:    []string{"Origin"},
    ExposeHeaders:   []string{"Content-Length"},
  }))

  log.Print("Creating timer to update data daily")
  go Run(loadData)

  log.Print("Running server")
  err := r.Run()
  Check(err)
}

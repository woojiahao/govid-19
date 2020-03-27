package server

import (
  "fmt"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/api"
  "github.com/woojiahao/govid-19/pkg/data"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "os"
)

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
  fmt.Println("Starting server...")
  // Use default server configurations
  r := gin.Default()

  // Run the data loading in a goroutine
  go loadData()

  // Load the API endpoints
  api.Build(r)

  // Auto-update the API data
  // TODO Test that this works
  Run(loadData)

  // Configure CORS for the API to allow all origins
  r.Use(cors.New(cors.Config{
    AllowAllOrigins: true,
    AllowMethods:    []string{"GET"},
    AllowHeaders:    []string{"Origin"},
    ExposeHeaders:   []string{"Content-Length"},
  }))

  // Run the server
  err := r.Run()
  fmt.Println("Server started.")
  Check(err)
}

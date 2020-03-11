package server

import (
  "fmt"
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
  r := gin.Default()
  go loadData()
  api.Build(r)
  err := r.Run()
  fmt.Println("Server started.")
  Check(err)
}

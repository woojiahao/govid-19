package main

import (
  "github.com/woojiahao/govid-19/pkg/api"
  "github.com/woojiahao/govid-19/pkg/server"
)

func main() {
  engine := server.Start()
  api.Build(engine)
}

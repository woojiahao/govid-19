package main

import (
  "github.com/woojiahao/govid-19/pkg/data"
  "github.com/woojiahao/govid-19/pkg/database"
  "log"
)

func main() {
  log.Print("Loading data")
  data.LoadData()

  log.Print("Setting up database")
  database.Setup(data.ConfirmedCases, data.RecoveredCases, data.DeathCases)

  //log.Print("Starting server")
  //server.Start()
}

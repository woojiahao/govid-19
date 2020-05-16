package main

import (
  "github.com/woojiahao/govid-19/pkg/data"
  "github.com/woojiahao/govid-19/pkg/database"
  "log"
)

var databaseManager *database.Manager

func main() {
  log.Print("Loading data")
  data.LoadData()

  log.Print("Setting up database")
  databaseManager = database.Setup(data.ConfirmedCases, data.RecoveredCases, data.DeathCases)

  log.Print("Starting server")
  //server.Start(databaseManager)
}

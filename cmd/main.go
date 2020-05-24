package main

import (
  "github.com/woojiahao/govid-19/pkg/data"
  "github.com/woojiahao/govid-19/pkg/database"
  "github.com/woojiahao/govid-19/pkg/server"
  "log"
)

func main() {
  log.Print("Connecting to database")
  databaseManager := database.Setup()

  if !databaseManager.IsUpToDate() {
    log.Print("Sources are out of date")
    log.Print("Updating sources")
    data.Update()

    log.Print("Loading data")
    confirmedCases, recoveredCases, deathCases := data.Load()

    log.Print("Uploading data into database")
    databaseManager.UploadData(confirmedCases, recoveredCases, deathCases)
  }

  log.Print("Starting server")
  server.Start(databaseManager)
}

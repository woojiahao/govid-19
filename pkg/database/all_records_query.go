package database

import (
  "fmt"
  "time"
)

type (
  displayType string
  sortOrder   string
  allRecords  struct {
    Date           time.Time
    Confirmed      int
    Recovered      int
    Deaths         int
    Country        string
    State          string
    Long           float32
    Lat            float32
    ConfirmedTotal int
    DeathsTotal    int
    RecoveredTotal int
  }
  AllRecordsQueryParams struct {
    first       int
    last        int
    country     string
    state       string
    sortTotal   sortOrder
    sortRecords sortOrder
    displayType
  }
)

const (
  Ascending        sortOrder   = "asc"
  Descending       sortOrder   = "desc"
  DisplayCountries displayType = "countries"
  DisplayStates    displayType = "states"
)

func (manager *Manager) queryAllRecords(params *AllRecordsQueryParams) []allRecords {
  totalSubQuery := manager.
    DB.
    Table("record").
    Select("location_id, sum(confirmed) as confirmed_total, sum(recovered) as recovered_total, sum(deaths) as deaths_total").
    Group("location_id").
    SubQuery()

  columnNames := []string{
    "r.date",
    "r.confirmed",
    "r.recovered",
    "r.deaths",
    "l.country",
    "l.state",
    "l.long",
    "l.lat",
    "s.confirmed_total",
    "s.deaths_total",
    "s.recovered_total",
  }
  query := manager.
    DB.
    Table("record r").
    Select(columnNames).
    Joins("inner join location l on l.id = r.location_id").
    Joins("inner join ? s on s.location_id = r.location_id", totalSubQuery)

  if params.country != "" {
    query = query.Where("l.country like ?", fmt.Sprintf("%%%s%%", params.country))
  }

  if params.state != "" {
    query = query.Where("l.state like ?", fmt.Sprintf("%%%s%%", params.state))
  }

  if params.sortTotal != "" {
    query = query.Order(fmt.Sprintf("confirmed_total %s", params.sortTotal))
  }

  if params.sortRecords != "" {
    query = query.Order("l.id").Order(fmt.Sprintf("confirmed %s", params.sortRecords))
  }

  var records []allRecords
  query.Find(&records)
  return records
}

func (manager *Manager) QueryAllRecords(params *AllRecordsQueryParams) {
  //allRecords := manager.queryAllRecords(params)
  switch params.displayType {
  case DisplayCountries:
    // Display only the general country information
    break
  case DisplayStates:
    // Display the statistics for all the states
    break
  default:
    panic("Invalid display type")
  }
}

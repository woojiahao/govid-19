package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "strconv"
  "time"
)

type (
  locationOverviewRecords struct {
    Date           time.Time
    Confirmed      int
    Recovered      int
    Deaths         int
    Id             uint
    Country        string
    State          string
    Long           float32
    Lat            float32
    ConfirmedTotal int
    DeathsTotal    int
    RecoveredTotal int
  }
  locationOverview struct {
    Id             uint              `json:"id"`
    Country        string            `json:"country"`
    State          string            `json:"state"`
    Long           float32           `json:"long"`
    Lat            float32           `json:"lat"`
    ConfirmedTotal int               `json:"confirmed_total"`
    RecoveredTotal int               `json:"recovered_total"`
    DeathsTotal    int               `json:"deaths_total"`
    Records        map[string]record `json:"records"`
  }
  record struct {
    Confirmed int `json:"confirmed"`
    Recovered int `json:"recovered"`
    Deaths    int `json:"deaths"`
  }
)

func hasLocation(locationId int) bool {
  var results []interface{}
  databaseManager.DB.Table("location").Where("id = ?", locationId).Find(&results)
  return len(results) > 0
}

func getLocationOverview(c *gin.Context) {
  locationId, err := strconv.Atoi(c.Param("location_id"))
  if err != nil {
    panic(newError(err, 500, "location_id must be an int"))
  }

  if !hasLocation(locationId) {
    panic(newError(nil, 404, "location_id does not match existing locations. Use /countries to get each location's id"))
  }

  columns := []string{
    "location_id",
    "sum(confirmed) confirmed_total",
    "sum(recovered) recovered_total",
    "sum(deaths) deaths_total",
  }
  subQuery := databaseManager.DB.Table("record").Select(columns).Group("location_id").SubQuery()

  columnNames := []string{
    "r.date",
    "r.confirmed",
    "r.recovered",
    "r.deaths",
    "l.id",
    "l.country",
    "l.state",
    "l.long",
    "l.lat",
    "s.confirmed_total",
    "s.deaths_total",
    "s.recovered_total",
  }
  query := databaseManager.
    DB.
    Table("record r").
    Select(columnNames).
    Joins("inner join location l on l.id = r.location_id").
    Joins("inner join ? s on s.location_id = r.location_id", subQuery).
    Where("l.id = ?", locationId)

  var records []locationOverviewRecords
  query.Find(&records)

  overview := locationOverview{
    Id:             records[0].Id,
    Country:        records[0].Country,
    State:          records[0].State,
    Long:           records[0].Long,
    Lat:            records[0].Lat,
    ConfirmedTotal: records[0].ConfirmedTotal,
    RecoveredTotal: records[0].RecoveredTotal,
    DeathsTotal:    records[0].DeathsTotal,
    Records:        make(map[string]record),
  }
  for _, r := range records {
    date := fmt.Sprintf("%02d-%02d-%d", r.Date.Day(), r.Date.Month(), r.Date.Year())
    overview.Records[date] = record{r.Confirmed, r.Recovered, r.Deaths}
  }

  OK(c, overview)
}

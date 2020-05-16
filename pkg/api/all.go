package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/data"
  "strconv"
  "time"
)

type response struct {
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

// TODO Test for case sensitivity in the query parameters
// TODO Add the sum of the data returned
func All(c *gin.Context) {
  params := c.Request.URL.Query()
  country, state, sortTotal, sortRecords, first, last := params.Get("country"),
    params.Get("state"),
    params.Get("sort-total"),
    params.Get("sort-records"),
    params.Get("first"),
    params.Get("last")

  totalSubquery := databaseManager.
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
  query := databaseManager.
    DB.
    Table("record r").
    Select(columnNames).
    Joins("inner join location l on l.id = r.location_id").
    Joins("inner join ? s on s.location_id = r.location_id", totalSubquery)

  if country != "" {
    query = query.Where("l.country like ?", fmt.Sprintf("%%%s%%", country))
  }

  if state != "" {
    query = query.Where("l.state like ?", fmt.Sprintf("%%%s%%", state))
  }

  if sortTotal != "" {
    order, status, errMsg := data.CheckSortOrder(sortTotal)
    if !status {
      BadRequest(c, errMsg)
      return
    }
    query = query.Order(fmt.Sprintf("confirmed_total %s", order))
  }

  if sortRecords != "" {
    order, status, errMsg := data.CheckSortOrder(sortRecords)
    if !status {
      BadRequest(c, errMsg)
      return
    }
    query = query.Order("l.id").Order(fmt.Sprintf("confirmed %s", order))
  }

  if first != "" {
    query = query.Limit(first)
  } else if last != "" {
    var count int
    query.Count(&count)
    fmt.Println(count)
    if n, err := strconv.Atoi(last); err != nil {
      query = query.Offset(count - n)
    }
  }

  result := make([]response, 0)
  query.Find(&result)
  fmt.Println(result[0])

  OK(c, result[0])

  //allSeries := data.NewAllSeries(confirmed, deaths, recovered)
  //OK(c, allSeries.ToJSON())
}

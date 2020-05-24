package api

import (
  "github.com/gin-gonic/gin"
)

// TODO Test for case sensitivity in the query parameters
func All(c *gin.Context) {
  //params := c.Request.URL.Query()
  //country, state, sortTotal, sortRecords, first, last := params.Get("country"),
  //  params.Get("state"),
  //  params.Get("sort-total"),
  //  params.Get("sort-records"),
  //  params.Get("first"),
  //  params.Get("last")
  //
  //databaseManager.QueryAllRecords(&database.AllRecordsQueryParams{
  //
  //})
  //
  //databaseResponses := make([]databaseResponse, 0)
  //query.Find(&databaseResponses)
  //
  //for _, r := range databaseResponses {
  //  fmt.Println(r)
  //}
  //
  //type categoryInformation struct {
  //  Total int
  //  Data  map[string]int
  //}
  //type stateInformation struct {
  //  Long      float32
  //  Lat       float32
  //  Confirmed categoryInformation
  //  Recovered categoryInformation
  //  Deaths    categoryInformation
  //}
  //data := make(map[string]map[string]stateInformation)
  //for _, record := range databaseResponses {
  //  if data[record.Country] == nil {
  //    data[record.Country] = make(map[string]stateInformation)
  //  }
  //
  //  data[record.Country][record.State] = stateInformation{
  //    record.Long,
  //    record.Lat,
  //    categoryInformation{
  //      Data: make(map[string]int),
  //    },
  //    categoryInformation{
  //      Data: make(map[string]int),
  //    },
  //    categoryInformation{
  //      Data: make(map[string]int),
  //    },
  //  }
  //  confirmed := data[record.Country][record.State].Confirmed
  //  confirmed.Data[record.Date.String()] = record.Confirmed
  //  confirmed.Total = record.ConfirmedTotal
  //
  //  recovered := data[record.Country][record.State].Recovered
  //  recovered.Data[record.Date.String()] = record.Recovered
  //  recovered.Total = record.RecoveredTotal
  //
  //  deaths := data[record.Country][record.State].Deaths
  //  deaths.Data[record.Date.String()] = record.Deaths
  //  deaths.Total = record.DeathsTotal
  //}
  //
  //OK(c, data)

  //if first != "" {
  //  query = query.Limit(first)
  //} else if last != "" {
  //  var count int
  //  query.Count(&count)
  //  fmt.Println(count)
  //  if n, err := strconv.Atoi(last); err != nil {
  //    query = query.Offset(count - n)
  //  }
  //}
  //
  //OK(c, databaseResponses)
}

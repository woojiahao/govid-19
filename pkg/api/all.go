package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/data"
)

// TODO Test for case sensitivity in the query parameters
// TODO Add the sum of the data returned
func All(c *gin.Context) {
  params := c.Request.URL.Query()
  //country, state, first, last, sortData, sortRecords := params.Get("country"),
  // params.Get("state"),
  // params.Get("first"),
  // params.Get("last"),
  // params.Get("sort-data"),
  // params.Get("sort-records")
  country, state, sortTotal := params.Get("country"), params.Get("state"), params.Get("sort-total")

  query := databaseManager.
    DB.
    Table("record").
    Select("*").
    Joins("inner join location on location.id = record.location_id")

  if country != "" {
    query = query.Where("location.country like ?", fmt.Sprintf("%%%s%%", country))
  }

  if state != "" {
    query = query.Where("location.state like ?", fmt.Sprintf("%%%s%%", state))
  }

  //// TODO Experiment with passing the sorting function as a lambda/function argument to clean up the code
  if sortTotal != "" {
   order, status, errMsg := data.CheckSortOrder(sortTotal)
   if !status {
     BadRequest(c, errMsg)
     return
   }
   fmt.Println(order)
  }

  //if sortRecords != "" {
  //  order, status, errMsg := data.CheckSortOrder(sortRecords)
  //  if !status {
  //    BadRequest(c, errMsg)
  //    return
  //  }
  //  confirmed, deaths, recovered = confirmed.SortRecords(order),
  //    deaths.SortRecords(order),
  //    recovered.SortRecords(order)
  //}
  //
  //if first != "" {
  //  num, status, errMsg := utility.CheckInt(first, "first", 0, len(confirmed.Records))
  //  if !status {
  //    BadRequest(c, errMsg)
  //    return
  //  }
  //  confirmed, deaths, recovered = confirmed.First(num),
  //    deaths.First(num),
  //    recovered.First(num)
  //} else if last != "" {
  //  num, status, errMsg := utility.CheckInt(last, "last", 0, len(confirmed.Records))
  //  if !status {
  //    BadRequest(c, errMsg)
  //    return
  //  }
  //  confirmed, deaths, recovered = confirmed.Last(num),
  //    deaths.Last(num),
  //    recovered.Last(num)
  //}
  //
  //allSeries := data.NewAllSeries(confirmed, deaths, recovered)
  //OK(c, allSeries.ToJSON())
}

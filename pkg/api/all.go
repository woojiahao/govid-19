package api

import (
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/data"
  "github.com/woojiahao/govid-19/pkg/utility"
)

// TODO Test for case sensitivity in the query parameters
// TODO Add the sum of the data returned
func All(c *gin.Context) {
  confirmed, deaths, recovered := data.GetAll()

  params := c.Request.URL.Query()
  country, state, first, last, sortData, sortRecords := params.Get("country"),
    params.Get("state"),
    params.Get("first"),
    params.Get("last"),
    params.Get("sort-data"),
    params.Get("sort-records")

  if country != "" {
    confirmed, deaths, recovered = confirmed.GetByCountry(country),
      deaths.GetByCountry(country),
      recovered.GetByCountry(country)
  }

  if state != "" {
    confirmed, deaths, recovered = confirmed.GetByState(state),
      deaths.GetByState(state),
      recovered.GetByState(state)
  }

  // TODO Experiment with passing the sorting function as a lambda/function argument to clean up the code
  if sortData != "" {
    order, status, errMsg := data.CheckSortOrder(sortData)
    if !status {
      BadRequest(c, errMsg)
      return
    }
    confirmed, deaths, recovered = confirmed.SortRecords(order),
      deaths.SortRecords(order),
      recovered.SortRecords(order)
  }

  if sortRecords != "" {
    order, status, errMsg := data.CheckSortOrder(sortRecords)
    if !status {
      BadRequest(c, errMsg)
      return
    }
    confirmed, deaths, recovered = confirmed.SortRecords(order),
      deaths.SortRecords(order),
      recovered.SortRecords(order)
  }

  if first != "" {
    num, status, errMsg := utility.CheckInt(first, "first", 0, len(confirmed.Records))
    if !status {
      BadRequest(c, errMsg)
      return
    }
    confirmed, deaths, recovered = confirmed.First(num),
      deaths.First(num),
      recovered.First(num)
  } else if last != "" {
    num, status, errMsg := utility.CheckInt(last, "last", 0, len(confirmed.Records))
    if !status {
      BadRequest(c, errMsg)
      return
    }
    confirmed, deaths, recovered = confirmed.Last(num),
      deaths.Last(num),
      recovered.Last(num)
  }

  allSeries := data.NewAllSeries(confirmed, deaths, recovered)
  OK(c, allSeries.ToJSON())
}

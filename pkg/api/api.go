package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/data"
  . "github.com/woojiahao/govid-19/pkg/utility"
)

var endpoints = []Endpoint{
  {
    RequestType: GET,
    Path:        "/ping",
    Action:      ping,
  },
  {
    RequestType: GET,
    Path:        "/all",
    Action:      all,
  },
}

func ping(c *gin.Context) {
  OK(c, gin.H{
    "message": "pong",
  })
}

// TODO Test for case sensitivity in the query parameters
// TODO Add the sum of the data returned
func all(c *gin.Context) {
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
    order, status, errMsg := checkSortOrder(sortData)
    if !status {
      BadRequest(c, errMsg)
      return
    }
    confirmed, deaths, recovered = confirmed.SortRecords(order),
      deaths.SortRecords(order),
      recovered.SortRecords(order)
  }

  if sortRecords != "" {
    order, status, errMsg := checkSortOrder(sortRecords)
    if !status {
      BadRequest(c, errMsg)
      return
    }
    confirmed, deaths, recovered = confirmed.SortRecords(order),
      deaths.SortRecords(order),
      recovered.SortRecords(order)
  }

  if first != "" {
    num, status, errMsg := CheckInt(first, "first", 0, len(confirmed.Records))
    if !status {
      BadRequest(c, errMsg)
      return
    }
    confirmed, deaths, recovered = confirmed.First(num),
      deaths.First(num),
      recovered.First(num)
  } else if last != "" {
    num, status, errMsg := CheckInt(last, "last", 0, len(confirmed.Records))
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

// TODO Move somewhere else
func checkSortOrder(raw string) (order data.SortOrder, status bool, errMsg string) {
  // TODO Better way to check against the constants
  switch raw {
  case "asc":
    order = data.Ascending
    status = true
  case "desc":
    order = data.Descending
    status = true
  default:
    order = ""
    status = false
    errMsg = "Invalid sort order. Available values are [ asc, desc ]"
  }
  return
}

func Build(engine *gin.Engine) {
  fmt.Println("Building endpoints...")
  for _, endpoint := range endpoints {
    path, action := endpoint.Path, endpoint.Action
    switch endpoint.RequestType {
    case GET:
      engine.GET(path, action)
    case POST:
      engine.POST(path, action)
    case PUT:
      engine.PUT(path, action)
    case DELETE:
      engine.DELETE(path, action)
    }
  }
}

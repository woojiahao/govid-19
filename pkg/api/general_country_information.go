package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "strings"
)

type (
  gciSortCol                string
  generalCountryInformation struct {
    Country   string `json:"country"`
    Confirmed int32  `json:"confirmed"`
    Recovered int32  `json:"recovered"`
    Deaths    int32  `json:"deaths"`
  }
  gciQueryParams struct {
    First int        `form:"first"`
    Last  int        `form:"last"`
    Sort  gciSortCol `form:"sort"`
    Order sortOrder  `form:"order"`
  }
)

const (
  confirmed gciSortCol = "confirmed"
  recovered gciSortCol = "recovered"
  deaths    gciSortCol = "deaths"
)

func (s *gciSortCol) isValid() bool {
  target := gciSortCol(strings.ToLower(string(*s)))

  switch target {
  case confirmed:
    fallthrough
  case recovered:
    fallthrough
  case deaths:
    return true
  default:
    return false
  }
}

// Returns the general information (non-specific) for every country in the world
// and the relevant statistics up till that day
func getGeneralCountryInformation(c *gin.Context) {
  var params gciQueryParams
  _ = c.Bind(&params)

  if !params.Sort.isValid() {
    params.Sort = confirmed
  }

  if !params.Order.isValid() {
    params.Order = descending
  }

  columns := []string{
    "l.country",
    "sum(r.confirmed) confirmed",
    "sum(r.recovered) recovered",
    "sum(r.deaths) deaths",
  }
  query := databaseManager.
    DB.
    Table("location l").
    Select(columns).
    Joins("inner join record r on l.id = r.location_id").
    Group("l.country").
    Order(fmt.Sprintf("%s %s", string(params.Sort), string(params.Order)))

  var results []generalCountryInformation
  query.Find(&results)

  if params.First != 0 {
    OK(c, results[:params.First])
  } else if params.Last != 0 {
    OK(c, results[len(results)-params.Last:])
  } else {
    OK(c, results)
  }
}

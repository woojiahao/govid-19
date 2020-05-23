package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "strconv"
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
  params := params(c, "first", "last", "sort", "order")
  first, last, sort, order := params[0], params[1], gciSortCol(params[2]), sortOrder(params[3])

  if !sort.isValid() {
    sort = confirmed
  }

  if !order.isValid() {
    order = descending
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
    Order(fmt.Sprintf("%s %s", string(sort), string(order)))

  var results []generalCountryInformation
  query.Find(&results)

  if f, err := strconv.Atoi(first); err == nil && f != 0 {
    OK(c, results[:f])
  } else if l, err := strconv.Atoi(last); err == nil && l != 0 {
    OK(c, results[len(results)-l:])
  } else {
    OK(c, results)
  }
}

package api

import (
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/database"
)

func getCountries(c *gin.Context) {
  columnNames := []string{
    "country",
    "state",
    "long",
    "lat",
  }
  query := databaseManager.
    DB.
    Table("location").
    Select(columnNames).
    Order("country").
    Order("state")

  var rawLocations []database.Location
  query.Find(&rawLocations)

  type location struct {
    State string  `json:"state,omitempty"`
    Long  float32 `json:"long"`
    Lat   float32 `json:"lat"`
  }
  locations := make(map[string][]location)
  for _, loc := range rawLocations {
    locations[loc.Country] = append(locations[loc.Country], location{loc.State, loc.Long, loc.Lat})
  }

  OK(c, locations)
}

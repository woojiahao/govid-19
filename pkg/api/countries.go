package api

import (
  "github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
  type databaseStructure struct {
    Country string
    State   string
    Long    float32
    Lat     float32
  }

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

  var databaseResponse []databaseStructure
  query.Find(&databaseResponse)

  type location struct {
    State string  `json:"state,omitempty"`
    Long  float32 `json:"long"`
    Lat   float32 `json:"lat"`
  }
  locations := make(map[string][]location)
  for _, loc := range databaseResponse {
    locations[loc.Country] = append(locations[loc.Country], location{loc.State, loc.Long, loc.Lat})
  }

  OK(c, locations)
}

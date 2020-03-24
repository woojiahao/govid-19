package api

import (
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/data"
)

// Returns the details of the countries with cases
func GetAvailableCountries(c *gin.Context) {
  countries := data.GetCountries()
  OK(c, countries)
}

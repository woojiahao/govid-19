package api

import "github.com/gin-gonic/gin"

func params(c *gin.Context, keys ...string) []string {
  var data []string
  params := c.Request.URL.Query()
  for _, key := range keys {
    data = append(data, params.Get(key))
  }
  return data
}

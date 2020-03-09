package api

import "github.com/gin-gonic/gin"

func OK(c *gin.Context, obj interface{}) {
  c.JSON(200, obj)
}

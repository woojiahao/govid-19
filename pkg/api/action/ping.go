package api

import (
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/api"
)

func Ping(c *gin.Context) {
  api.OK(c, gin.H{
    "message": "pong",
  })
}

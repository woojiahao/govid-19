package api

import (
  "github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
  OK(c, gin.H{
    "message": "pong",
  })
}

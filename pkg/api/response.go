package api

import (
  "github.com/gin-gonic/gin"
  "time"
)

type Error struct {
  status    int
  message   string
  timestamp int64
}

func NewError(status int, message string) *Error {
  return &Error{
    status:    status,
    message:   message,
    timestamp: time.Now().Unix(),
  }
}

func (e *Error) toJSON() map[string]interface{} {
  return gin.H{
    "status": e.status,
    "message": e.message,
    "timestamp": e.timestamp,
  }
}

func OK(c *gin.Context, obj interface{}) {
  c.JSON(200, obj)
}

func BadRequest(c *gin.Context, msg string) {
  c.JSON(400, NewError(400, msg).toJSON())
}

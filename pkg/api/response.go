package api

import (
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/utility"
  "time"
)

type Error struct {
  Status    int       `json:"status"`
  Message   string    `json:"message"`
  Err       error     `json:"error,omitempty"`
  Timestamp time.Time `json:"timestamp"`
}

func newError(err error, status int, message string) *Error {
  return &Error{status, message, err, time.Now()}
}

func (e *Error) Error() string {
  buf, err := json.Marshal(e)
  utility.Check(err)
  return string(buf)
}

func OK(c *gin.Context, obj interface{}) {
  c.JSON(200, obj)
}

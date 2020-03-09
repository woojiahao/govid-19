package api

import "github.com/gin-gonic/gin"

type RequestType string

const (
  GET    RequestType = "GET"
  POST   RequestType = "POST"
  PUT    RequestType = "PUT"
  DELETE RequestType = "DELETE"
)

type Endpoint struct {
  RequestType RequestType
  Path        string
  Action      func(c *gin.Context)
}

package server

import (
  "github.com/gin-gonic/gin"
  "github.com/woojiahao/govid-19/pkg/api"
)

func handleError() gin.HandlerFunc {
  return func(c *gin.Context) {
    defer func() {
      if err := recover(); err != nil {
        switch err.(type) {
        case *api.Error:
          apiError := err.(*api.Error)
          c.JSON(apiError.Status, apiError)
        }

        panic(err)
      }
    }()

    c.Next()
  }
}

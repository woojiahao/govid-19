package server

import (
	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	r := gin.Default()
	err := r.Run()
	if err != nil {
		panic(err)
	}
	return r
}

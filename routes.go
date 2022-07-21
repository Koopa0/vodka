package main

import (
	"github.com/gin-gonic/gin"
)

func routes() *gin.Engine {
	r := gin.Default()

	r.GET("/getUser", getUser)

	return r
}

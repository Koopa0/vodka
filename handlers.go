package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {

	fmt.Fprintf(c.Writer, "Hello world!")
}

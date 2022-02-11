package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloworld() string {
	return "Hello World!!"
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, helloworld())
	})

	r.Run()
}

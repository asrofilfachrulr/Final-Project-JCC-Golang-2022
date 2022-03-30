package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("Bismillah Final Project"))
	})

	r.Run()
}

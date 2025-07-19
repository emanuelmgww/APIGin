package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	rota := gin.Default()

	rota.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensagem": "Hello, World",
		})
	})

	rota.Run(":3333")
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	tst()
}

func tst() {
	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "home",
		})
	})

	engine.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})

	engine.Run()
}

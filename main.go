package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func main() {

	tst()
}

func tst() {
	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("home %d", rand.Int()),
		})
	})

	engine.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("ping %d", rand.Int()),
		})
	})

	engine.Run()
}

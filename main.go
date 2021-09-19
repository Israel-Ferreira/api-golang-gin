package main

import (
	"github.com/Israel-Ferreira/challenge-conductor/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/",func(c *gin.Context) {
		helloMsg := map[string] string {
			"msg": "Hello World",
		}

		c.JSON(200, helloMsg)
	})


	r.GET("/cards/:id", services.GetCardById)

	
	r.Run()
}


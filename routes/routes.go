package routes

import (
	"github.com/gin-gonic/gin"
)

var TablesInMemory []string

type Command struct {
	Query string `json:"query"`
}

func CreateRoutes() {

	TablesInMemory = []string{"aa", "ba"}

	router := gin.Default()

	router.GET("/table", func(c *gin.Context) {
		c.String(200, TablesInMemory[0])
	})

	router.POST("/command", commandRoute)

	router.Run()
}

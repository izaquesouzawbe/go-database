package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes() {

	router := gin.Default()

	router.GET("/tables", func(c *gin.Context) {
		c.String(200, "TablesInMemory[0]")
	})

	router.POST("/command", commandRoute)

	router.Run()
}

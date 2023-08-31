package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes() {

	router := gin.Default()

	router.POST("/command", routeCommand)
	router.POST("/insert", routeInsert)
	router.POST("/insertTableColumn", routeInsert)

	router.Run()
}

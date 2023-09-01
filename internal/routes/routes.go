package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes() {

	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	router.POST("/command", routeCommand)
	router.POST("/insert", routeInsert)

	router.Run()
}

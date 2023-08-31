package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-zdb-api/internal/models/request"
)

func routeInsert(c *gin.Context) {

	var insert request.InsertRequest
	err := c.BindJSON(&insert)
	if err != nil {
		return
	}

	fmt.Println(insert)

	//lines := commands.RunCommand(insert)

	c.JSON(200, "ok")
}

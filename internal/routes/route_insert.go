package routes

import (
	"github.com/gin-gonic/gin"
	"go-zdb-api/internal/commands/commands_list"
	"go-zdb-api/internal/models/request"
	"go-zdb-api/pkg/general"
)

func routeInsert(c *gin.Context) {

	general.RuntimeStarted()

	var insert request.InsertRequest
	err := c.BindJSON(&insert)
	if err != nil {
		return
	}

	commands_list.RouteInsertInto(insert)

	general.RuntimeDone()
	c.JSON(200, "ok")

}

package routes

import (
	"github.com/gin-gonic/gin"
	"go-zdb-api/internal/commands"
	"go-zdb-api/internal/models/request"
	"net/http"
)

func routeCommand(c *gin.Context) {

	var command request.RequestCommand
	err := c.BindJSON(&command)
	if err != nil {
		return
	}

	lines := commands.RunCommand(command.Value)

	c.String(http.StatusOK, string(lines))
	//c.JSON(200)
}

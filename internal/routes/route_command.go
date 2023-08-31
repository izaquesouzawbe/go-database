package routes

import (
	"github.com/gin-gonic/gin"
	"go-zdb-api/internal/commands"
	"go-zdb-api/internal/models"
)

func commandRoute(c *gin.Context) {

	var command models.Command
	err := c.BindJSON(&command)
	if err != nil {
		return
	}

	lines := commands.RunCommand(command.Value)

	c.JSON(200, lines)
}

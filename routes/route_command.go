package routes

import (
	"github.com/gin-gonic/gin"
	"go-database/commands"
	"log"
)

func commandRoute(c *gin.Context) {

	var command Command
	c.BindJSON(&command)

	log.Println(command)

	commands.RunCommand(command.Query)

	c.JSON(200, command)
}

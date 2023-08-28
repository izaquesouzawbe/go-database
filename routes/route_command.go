package routes

import (
	"github.com/gin-gonic/gin"
	"go-database/commands"
	"log"
)

func commandRoute(c *gin.Context) {

	var command Command
	err := c.BindJSON(&command)
	if err != nil {
		return
	}

	log.Println(command)

	lines := commands.RunCommand(command.Query)

	//fmt.Println(lines)

	c.JSON(200, lines)
}

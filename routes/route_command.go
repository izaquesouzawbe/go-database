package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-database/commands"
)

func commandRoute(c *gin.Context) {
	var command Command
	c.BindJSON(&command)

	fmt.Println(command)

	commands.RunCommand(command.Query)

	//command.CommandValue = cleanCommand(command.CommandValue)
	//comands := getListCommand(command.CommandValue)

	//fmt.Println(command)
	//fmt.Println(comands)

	//value := commandSelectTable(comands)

	c.JSON(200, command)
}

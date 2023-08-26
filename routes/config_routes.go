package routes

import (
	"github.com/gin-gonic/gin"
)

var tablesInMemory []string

type Command struct {
	CommandValue string `json:"command"`
}

func CreateRoutes() {

	tablesInMemory = []string{"aa", "ba"}

	router := gin.Default()

	router.GET("/table", func(c *gin.Context) {
		c.String(200, tablesInMemory[0])
	})

	router.POST("/command", func(c *gin.Context) {

		/*var command Command
		c.BindJSON(&command)

		command.CommandValue = cleanCommand(command.CommandValue)
		comands := getListCommand(command.CommandValue)

		fmt.Println(command)
		fmt.Println(comands)

		value := commandSelectTable(comands)

		c.JSON(200, value)*/
	})

	router.Run()
}

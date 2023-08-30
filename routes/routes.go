package routes

import (
	"github.com/gin-gonic/gin"
	"go-database/commands"
)

var TablesInMemory []commands.Table

type TypeCommand string

const (
	Database TypeCommand = "Database"
	Table    TypeCommand = "Table"
	Sequence TypeCommand = "Sequence"
	Insert   TypeCommand = "Insert"
)

type Command struct {
	Type  TypeCommand `json:"type"`
	Query string      `json:"query"`
}

func CreateRoutes() {

	router := gin.Default()

	router.GET("/tables", func(c *gin.Context) {
		c.String(200, "TablesInMemory[0]")
	})

	router.POST("/command", commandRoute)

	router.Run()
}

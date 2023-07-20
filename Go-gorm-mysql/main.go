package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed string `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "clean room", Completed: "false"},
	{ID: "2", Item: "eat lunch", Completed: "false"},
	{ID: "3", Item: "eat dinner", Completed: "false"},
}

func gettodos(context *gin.Context) {

	context.JSON(http.StatusOK, todos)
}
func addtodos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/", gettodos)
	router.POST("/", addtodos)
	router.Run("localhost:9090")
}

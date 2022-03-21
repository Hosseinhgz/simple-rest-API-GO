package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "01", Item: "Clean Room", Completed: false},
	{ID: "02", Item: "Self study GO", Completed: false},
	{ID: "03", Item: "Read Articles", Completed: false},
	{ID: "04", Item: "Email to the company", Completed: false},
}

// we are specifying the type "*gin.Context"
func getTodos(context *gin.Context) {

	// convert from todo type to JSON
	context.IndentedJSON(http.StatusOK, todos)
}
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo is not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo is not Found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

// second parantise is going to specify 2 returns - check that we get todo type or we going to have error - one of them is going to happen
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo with this id is not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todo", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)

	router.Run("localhost:9090")
}

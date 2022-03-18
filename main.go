package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID					string `json:"id"`
	Item				string `json:"title"`
	Completed		bool   `json:"completed"`
}

var todos = []todo{
	{ID:"01", Item:"Clean Room", Completed: false},
	{ID:"02", Item:"Self study GO", Completed: false},
	{ID:"03", Item:"Read Articles", Completed: false},
	{ID:"04", Item:"Email to the company", Completed: false},
}

func main(){
	router := gin.Default()
	
}
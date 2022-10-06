package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

var todos = []todo{
	{
		ID:    "1",
		Name:  "Code a todo app",
		Owner: "Owner 1",
	},
	{
		ID:    "2",
		Name:  "Play some games",
		Owner: "Owner 2",
	},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, todos)
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func updateTodoById(c *gin.Context) {
	id := c.Param("id")
	for index, todo := range todos {
		if todo.ID == id {
			var updatedTodo = todo
			if err := c.BindJSON(&updatedTodo); err != nil {
				return
			}
			todos[index] = updatedTodo
			c.IndentedJSON(http.StatusOK, todos)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func deleteTodoById(c *gin.Context) {
	id := c.Param("id")
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			c.IndentedJSON(http.StatusOK, todos)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/todos", getTodos)
	r.POST("/todos", addTodo)
	r.GET("/todos/:id", getTodoById)
	r.PUT("/todos/:id", updateTodoById)
	r.DELETE("/todos/:id", deleteTodoById)

	r.Run("localhost:8080")
}

package main

import (
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{}
var nextID int64 = 0

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, t := range todos {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil || newTodo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	newTodo.ID = atomic.AddInt64(&nextID, 1)
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updated Todo
	if err := c.ShouldBindJSON(&updated); err != nil || updated.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updated.Title
			todos[i].Done = updated.Done
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func main() {
	r := gin.Default()

	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodoByID)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run(":8080") // Start server
}

package main

import (
	"todolist/config"
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	r := gin.Default()

	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	
	r.Run(":8080")
}
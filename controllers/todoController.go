package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	fmt.Println("todoリストの取得")
}

func CreateTodo(c * gin.Context) {
	fmt.Println("todoリストの作成")
}
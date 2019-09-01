package main

import (
	"GOLANG/todo/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	task := new(handler.TodoController)
	r.POST("/create", task.Create)
	r.GET("/fetch", task.Get)
	r.GET("/completed", task.Completed)
	r.POST("/update", task.Update)
	r.POST("/delete", task.Delete)
	r.Run(":8888")
}

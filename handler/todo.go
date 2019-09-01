package handler

import (
	"GOLANG/todo/forms"
	"GOLANG/todo/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

var taskModel = new(models.TaskModel)

func (task *TodoController) Create(c *gin.Context) {
	var data forms.CreateTaskCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := taskModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Task could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Task Created", "form": data})
}

func (task *TodoController) Get(c *gin.Context) {

	tasks, err := taskModel.Get()
	if err != nil {
		c.JSON(406, gin.H{"message": "Task could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Tasks Fetched", "form": tasks})
}

func (task *TodoController) Completed(c *gin.Context) {

	tasks, err := taskModel.Completed()
	if err != nil {
		c.JSON(406, gin.H{"message": "Task could not be completed", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Tasks Completed", "form": tasks})
}

func (task *TodoController) Update(c *gin.Context) {
	//id := c.Param("id")
	data := forms.UpdateTaskCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}
	fmt.Println(data)
	err := taskModel.Update(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Task Could Not Be Updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Task Updated"})
}

func (task *TodoController) Delete(c *gin.Context) {
	data := forms.UpdateTaskCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}
	fmt.Println(data)
	err := taskModel.Delete(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Task Could Not Be Deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Task Deleted"})
}

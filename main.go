package main

import (
	"github.com/DuongWuangDat/todolist-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/tasks", routes.GetTasks)
	r.GET("/tasks/:id", routes.GetByID)
	r.POST("/tasks", routes.CreateTask)
	r.PUT("/tasks/:id", routes.UpdateTask)
	r.DELETE("/tasks/:id", routes.DeleteTask)

	r.Run(":8080")
}

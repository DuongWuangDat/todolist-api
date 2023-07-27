package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DuongWuangDat/todolist-api/database"
	"github.com/DuongWuangDat/todolist-api/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	query := c.DefaultQuery("title", "")
	var tasks []models.Task
	if query == "" {

		err := database.Db.Find(&tasks).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": tasks,
		})
		return
	}
	err := database.Db.Where("title = ?", query).Find(&tasks).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	err := database.Db.First(&task, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func CreateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	task.CreatedAt = time.Now().Unix()
	task.IsDone = false
	err = database.Db.Create(&task).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Created succesfully",
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var body models.Task
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var task models.Task
	err = database.Db.First(&task, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err = database.Db.Find(&task).Updates(models.Task{
		Title:     body.Title,
		IsDone:    body.IsDone,
		CreatedAt: time.Now().Unix(),
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated succesfully",
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	err := database.Db.First(&task, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err = database.Db.Delete(&task).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted succesfully",
	})
}

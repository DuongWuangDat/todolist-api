package database

import (
	"log"

	"github.com/DuongWuangDat/todolist-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	d, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connect sucessfully")
	Db = d
	Db.AutoMigrate(&models.Task{})

}

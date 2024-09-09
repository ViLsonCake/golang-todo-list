package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"todo-list-api/config"
	"todo-list-api/entity"
)

func main() {
	webConfig := config.GetWebConfig()
	router := gin.Default()

	var db *gorm.DB

	db, err := gorm.Open("postgres", config.GetPostgresUrl())
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Todo{})

	router.Run(":" + webConfig.Port)
}

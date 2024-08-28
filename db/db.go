package db

import (
	"user-service/config"
	"user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *config.Config) {

	database, err := gorm.Open(postgres.Open(config.DBConnectionString), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos")
	}

	database.AutoMigrate(&models.User{})

	DB = database
}

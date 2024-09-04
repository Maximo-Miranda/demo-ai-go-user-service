package db

import (
	"user-service/config"
	"user-service/models"

	"gorm.io/driver/postgres" // Driver de PostgreSQL para GORM
	"gorm.io/gorm"            // ORM para Go
)

// DB es una variable global que contiene la conexión a la base de datos
var DB *gorm.DB

// ConnectDatabase establece la conexión con la base de datos y configura el ORM
func ConnectDatabase(config *config.Config) {

	// Abre una conexión a la base de datos usando la cadena de conexión proporcionada
	database, err := gorm.Open(postgres.Open(config.DBConnectionString), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos")
	}

	// Realiza la migración automática de la estructura User a la base de datos
	database.AutoMigrate(&models.User{})

	// Asigna la conexión establecida a la variable global DB
	DB = database
}

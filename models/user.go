package models

import "gorm.io/gorm" // Importa GORM, un ORM para Go que facilita las operaciones de base de datos

// User define la estructura de datos para un usuario en la aplicación
type User struct {
	gorm.Model        // Incorpora los campos ID, CreatedAt, UpdatedAt y DeletedAt de GORM
	FirstName  string `json:"first_name"`           // Nombre del usuario
	LastName   string `json:"last_name"`            // Apellido del usuario
	Email      string `json:"email" gorm:"unique"`  // Correo electrónico (único en la base de datos)
	Password   string `json:"password,omitempty"`   // Contraseña (omitida en las respuestas JSON)
	Phone      string `json:"phone"`                // Número de teléfono
	Address    string `json:"address"`              // Dirección
	City       string `json:"city"`                 // Ciudad
	State      string `json:"state"`                // Estado o provincia
	Country    string `json:"country"`              // País
	UserToken  string `json:"user_token,omitempty"` // Token de usuario (omitido en las respuestas JSON)
}

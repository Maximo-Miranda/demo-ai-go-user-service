package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Importa la librería para cargar variables de entorno desde un archivo .env
)

// Config almacena la configuración de la aplicación
type Config struct {
	DBConnectionString string // Cadena de conexión a la base de datos
	APPEnv             string // Entorno de la aplicación (ej. development, production)
}

// LoadConfig carga la configuración desde variables de entorno o archivo .env
func (c *Config) LoadConfig() {

	// Si no estamos en producción, intentamos cargar el archivo .env
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No se encontró el archivo .env") // No es un error crítico, solo informativo
		}
	}

	// Cargamos las variables de entorno en la estructura Config
	c.DBConnectionString = os.Getenv("DB_CONNECTION_STRING")
	c.APPEnv = os.Getenv("APP_ENV")
}

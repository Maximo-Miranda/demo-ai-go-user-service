package main

import (
	"fmt"
	"time"
	"user-service/config"
	"user-service/handlers"
	"user-service/middleware"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	echo_middlewares "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Inicialización de la configuración
	conf := config.Config{}
	conf.LoadConfig()

	// Retraso de 5 segundos en entornos no de desarrollo
	if conf.APPEnv != "development" {
		time.Sleep(time.Second * 5)
	}

	// Conexión a la base de datos
	//db.ConnectDatabase(&conf)

	// Configuración del servidor Echo
	e := echo.New()
	e.Use(echo_middlewares.Logger())

	podName := fmt.Sprintf("user-service-%s", uuid.New().String())
	messageResponse := fmt.Sprintf("User service is running! v1.0.1 - %s", podName)

	// Ruta de prueba para verificar que el servicio está en funcionamiento
	e.GET("/", func(c echo.Context) error {
		return c.String(200, messageResponse)
	})

	// Rutas para las operaciones de usuarios
	e.POST("/register", handlers.Register) // Registro de nuevos usuarios
	e.POST("/login", handlers.Login)       // Inicio de sesión
	e.GET("/users", handlers.ListUsers)    // Listar todos los usuarios

	// Ruta protegida para validar el token de usuario
	e.GET("/validate", handlers.ValidateToken, middleware.ValidateToken)

	// Iniciar el servidor en el puerto 8080
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"time"
	"user-service/config"
	"user-service/db"
	"user-service/handlers"
	"user-service/middleware"

	"github.com/labstack/echo/v4"
	echo_middlewares "github.com/labstack/echo/v4/middleware"
)

func main() {

	conf := config.Config{}
	conf.LoadConfig()

	if conf.APPEnv != "development" {
		time.Sleep(time.Second * 5)
	}
	db.ConnectDatabase(&conf)

	e := echo.New()
	e.Use(echo_middlewares.Logger())

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
	e.GET("/users", handlers.ListUsers)

	// Nueva ruta protegida
	e.GET("/validate", handlers.ValidateToken, middleware.ValidateToken)

	e.Logger.Fatal(e.Start(":8080"))
}

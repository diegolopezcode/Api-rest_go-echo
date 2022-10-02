package main

import (
	"log"

	router "github.com/DiegoLopez-ing/api_rest/api/routers"
	"github.com/DiegoLopez-ing/api_rest/cmd/injectors"
	"github.com/DiegoLopez-ing/api_rest/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	serverConfig, err := configs.LoadServiceConfig("service_config.json")
	if err != nil {
		log.Fatalln("No se cargaorn las configuraciones de los servicios:", err.Error())
	}
	log.Println("Server OK!")
	if err := injectors.LoadConfig("db_config.json"); err != nil {
		log.Fatalln("No se cargaron las configuraciones de la base de datos", err.Error())
	}
	log.Println("Base de datos OK!")
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: serverConfig.Cors.AllowedOrigins,
		AllowMethods: serverConfig.Cors.AllowedMethods,
	}))
	server.HideBanner = true
	router.Upgrade(server)
	if err := server.Start(serverConfig.Addres); err != nil {
		log.Fatal("Error al iniciar el servidor", err.Error())
	}
	log.Println("Servidor iniciado correctamente")
}

package main
import (
	"log"
	 "github.com/labstack/echo/v4"
	 "github.com/labstack/echo/v4/middleware"
	 "github.com/DiegoLopez-ing/api_rest/cmd/injectors"
	 "github.com/DiegoLopez-ing/api_rest/configs"
	 "github.com/DiegoLopez-ing/api_rest/api/routers"

func main() {
	serverConfig, err := configs.LoadServerConfigs("service_config.json")
	if err != nil {
		log.Fatalln("No se cargaorn las configuraciones de los servicios")
	}
	log.PrintLog("Server OK!")
	if err := injectors.LoadConfig("db_config.json"); err != nil {
		log.Fatalln("No se cargaron las configuraciones de la base de datos")
	}
	logs.PrintLog("Base de datos OK!")
	server := echo.New()
	server.Use(middleware.CORSWithConfigs(middleware.CORSConfig{
		AllowOrigins: serverConfig.Cors.AllowOrigins,
		AllowMethods: serverConfig.Cors.AllowMethods,
	}))
	server.HideBanner = true
	router.Upgrade(server)
err:= server.Start(serverConfig.Server.Address)
if err != nil {
	logs.Fatal("Error al iniciar el servidor")
}
log.Println("Servidor iniciado correctamente")
}

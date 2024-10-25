package api

import (
	"enhanced-notes/config"
	"enhanced-notes/src/api/rest"
	"enhanced-notes/src/api/rest/handler"
	"enhanced-notes/src/helper"

	"github.com/gofiber/fiber/v3"
)


func StartServer(config config.AppConfig){
	app := fiber.New();

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App: app,
		Auth: auth,
		Config: config,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)

}	

func setupRoutes(rh *rest.RestHandler) {
	handler.SetupUserRoutes()
}
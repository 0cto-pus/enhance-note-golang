package rest

import (
	"enhanced-notes/config"
	"enhanced-notes/src/helper"

	"github.com/gofiber/fiber/v3"
)

type RestHandler struct {
	App *fiber.App
	//DB *gorm.DB
	Auth helper.Auth
	Config config.AppConfig
}
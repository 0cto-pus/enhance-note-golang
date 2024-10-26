package api

import (
	"enhanced-notes/config"
	"enhanced-notes/src/api/rest"
	"enhanced-notes/src/api/rest/handler"
	"enhanced-notes/src/domain"
	"enhanced-notes/src/helper"
	"log"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func StartServer(config config.AppConfig){
	app := fiber.New();
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{});

	if err != nil{
		log.Fatalf("database connection error %v\n", err)
	}

	log.Printf("database connected")

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("error on running migration %v", err.Error()) 
	}
	log.Println("migration successful") 

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App: app,
		DB: db,
		Auth: auth,
		Config: config,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)

}	

func setupRoutes(rh *rest.RestHandler) {
	handler.SetupUserRoutes()
}

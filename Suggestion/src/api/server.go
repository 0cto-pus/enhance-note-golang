package api

import (
	"enhance-notes-suggestion/config"
	"enhance-notes-suggestion/src/api/rest"
	"enhance-notes-suggestion/src/api/rest/handler"
	"enhance-notes-suggestion/src/domain"
	"enhance-notes-suggestion/src/helper"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

	err = db.AutoMigrate(&domain.Suggestion{})
	if err != nil {
		log.Fatalf("error on running migration %v", err.Error())
	}
	log.Println("migration successful")

	// cors configuration
	c := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:9003"},
		AllowHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	app.Use(c)

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
	//Note Routes
	handler.SetupNoteRoutes(rh)
}

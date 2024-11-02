package api

/* import (
	"enhance-notes-note-service/config"
	"enhance-notes-note-service/src/api/rest"
	"enhance-notes-note-service/src/api/rest/handler"
	"enhance-notes-note-service/src/domain"
	"enhance-notes-note-service/src/helper"
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

	err = db.AutoMigrate(&domain.Note{})
	if err != nil {
		log.Fatalf("error on running migration %v", err.Error())
	}
	log.Println("migration successful")

	// cors configuration
	c := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:9001"},
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
	//User Routes
	handler.SetupNoteRoutes(rh)
}
*/
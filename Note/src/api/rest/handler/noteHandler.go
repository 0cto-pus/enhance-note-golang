package handler

/* import (
	"enhance-notes-note-service/src/api/rest"
	"enhance-notes-note-service/src/repository"
	"enhance-notes-note-service/src/service"

	"github.com/gofiber/fiber/v3"
)

type NoteController struct {
	noteService service.INoteService
}

func NewNoteController(service service.INoteService) *NoteController{
	return &NoteController{
		noteService: service,
	}
}

func SetupNoteRoutes(rh *rest.RestHandler) {
	app:=rh.App
	service := service.NewNoteService(repository.NewNoteRepository(rh.DB), rh.Auth, rh.Config)

	handler := NoteController{
		noteService: service,
	}

	pubRoutes := app.Group("/users")
	//Public endpoint

	pubRoutes.Post("/signup", handler.SignUp)
	pubRoutes.Post("/login", handler.Login)
}

func (handler *NoteController) SignUp(ctx fiber.Ctx) error {
	return nil
}

func (handler *NoteController) Login(ctx fiber.Ctx) error {
	return nil
} */
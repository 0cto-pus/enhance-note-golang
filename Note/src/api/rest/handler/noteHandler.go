package handler

import (
	"enhance-notes-note-service/src/api/rest"
	"enhance-notes-note-service/src/dto"
	"enhance-notes-note-service/src/helper"
	"enhance-notes-note-service/src/repository"
	"enhance-notes-note-service/src/service"

	"github.com/gofiber/fiber/v3"
)

type NoteController struct {
	noteService service.INoteService
	auth helper.Auth
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
		auth: rh.Auth,
	}

	pubRoutes := app.Group("/notes",rh.Auth.Authorize)
	//Public endpoint

	pubRoutes.Post("/createnote" ,handler.CreateNote)
	pubRoutes.Get("/getnotes", handler.GetUserNotes)
}

func (handler *NoteController) CreateNote(ctx fiber.Ctx) error {
	userId := handler.auth.GetCurrentUserID(ctx)
    if userId == 0 {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "unauthorized",
        })
    }
	userInput:= dto.NoteCreate{}

    if err := ctx.Bind().Body(&userInput); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "please provide valid input",
        })
    }
    _, err := handler.noteService.CreateNote(userInput, userId)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "failed to create note",
        })
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Note added",
    })
}

func (handler *NoteController) GetUserNotes(ctx fiber.Ctx) error {
	userId := handler.auth.GetCurrentUserID(ctx) //We get user_id from auth token.
    if userId == 0 {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "unauthorized",
        })
    }
	notes, err := handler.noteService.GetUserNotes(userId)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "failed to retrieve notes",
        })
    }

  
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "notes": notes,
    })
} 
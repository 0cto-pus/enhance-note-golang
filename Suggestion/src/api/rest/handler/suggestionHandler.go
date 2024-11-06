package handler

import (
	"enhance-notes-suggestion/src/api/rest"
	"enhance-notes-suggestion/src/dto"
	"enhance-notes-suggestion/src/helper"
	"enhance-notes-suggestion/src/repository"
	"enhance-notes-suggestion/src/service"

	"github.com/gofiber/fiber/v3"
)

type SuggestionController struct {
	suggestionService service.ISuggestionService
	auth helper.Auth
}

func NewSuggestionController(service service.ISuggestionService) *SuggestionController{
	return &SuggestionController{
		suggestionService: service,
	}
}

func SetupNoteRoutes(rh *rest.RestHandler) {
	app:=rh.App
	service := service.NewSuggestionService(repository.NewSuggestionRepository(rh.DB), rh.Auth, rh.Config)

	handler := SuggestionController{
		suggestionService: service,
		auth: rh.Auth,
	}

	pubRoutes := app.Group("/suggestions",rh.Auth.Authorize)
	//Public endpoint

	pubRoutes.Post("/createsuggestion" ,handler.CreateSuggestion)
	pubRoutes.Get("/getsuggestions", handler.GetUserNotes)
}

func (handler *SuggestionController) CreateSuggestion(ctx fiber.Ctx) error {
	userId := handler.auth.GetCurrentUserID(ctx)
    if userId == 0 {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "unauthorized",
        })
    }
	userInput:= dto.SuggestioneCreate{}

    if err := ctx.Bind().Body(&userInput); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "please provide valid input",
        })
    }
    _, err := handler.suggestionService.CreateSuggestion(userInput, userId)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "failed to create suggestion",
        })
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Suggestion added",
    })
}

func (handler *SuggestionController) GetUserNotes(ctx fiber.Ctx) error {
	userId := handler.auth.GetCurrentUserID(ctx) //We get user_id from auth token.
    if userId == 0 {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "unauthorized",
        })
    }
	suggestions, err := handler.suggestionService.GetUserSuggestions(userId)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "failed to retrieve suggestions",
        })
    }


    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "suggestion": suggestions,
    })
}  
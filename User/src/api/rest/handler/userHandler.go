package handler

import (
	"enhanced-notes/src/api/rest"
	"enhanced-notes/src/dto"
	"enhanced-notes/src/repository"
	"enhanced-notes/src/service"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(service service.IUserService) *UserController{
	return &UserController{
		userService: service,
	}
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app:=rh.App
	service := service.NewUserService(repository.NewUserRepository(rh.DB), rh.Auth, rh.Config)

	handler := UserController{
		userService: service,
	}

	pubRoutes := app.Group("/users")
	//Public endpoint

	pubRoutes.Post("/signup", handler.SignUp)
	pubRoutes.Post("/login", handler.Login)
}

func (handler *UserController) SignUp(ctx fiber.Ctx) error {
	var err error
	var token string
	userInput:= dto.UserSignUp{}

	err= ctx.Bind().Body(&userInput)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "please provide valid input",
		})
	}

	token, err = handler.userService.SignUp(userInput)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "error on sign up",
		})
	}


	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "register",
		"token" : token,
	})
}

func (handler *UserController) Login(ctx fiber.Ctx) error {
	var err error
	var token string
	userInput:= dto.UserLogin{}

	err= ctx.Bind().Body(&userInput)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "please provide valid input",
		})
	}

	token, err = handler.userService.Login(userInput)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "please provide correct user id password",
		})
	}


	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login",
		"token" : token,
	})
}
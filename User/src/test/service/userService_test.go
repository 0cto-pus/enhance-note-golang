package service

import (
	"enhanced-notes/src/domain"
	"enhanced-notes/src/dto"
	"enhanced-notes/src/service"
	"os"
	"testing"
)

var userService service.IUserService

func TestMain(m *testing.M){
	initialUsers := []domain.User{
		{
		ID: 1,
		Email: "test@test.com",
		Password: "testing123app",
		},
		{
		ID: 2,
		Email: "test1@test.com",
		Password: "testing123app",
		},
	}
	mockRepository := NewMockUserRepository(initialUsers)
	userService =service.NewUserService(mockRepository)
	exitCode:=m.Run()
	os.Exit(exitCode)
}


func Test_ShouldCreateUserAndPassToken(t *testing.T){
	userService.SignUp(dto.UserSignUp{UserLogin: dto.UserLogin{Email:"test3@test.com", Password:"testtest123123"}})
	t.Run("ShouldCreateUserAndPassToken", func(t *testing.T) {

	})
}
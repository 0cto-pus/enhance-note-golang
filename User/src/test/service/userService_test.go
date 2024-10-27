package service

import (
	"enhanced-notes/src/domain"
	"enhanced-notes/src/dto"
	"enhanced-notes/src/helper"
	"enhanced-notes/src/service"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var userService service.IUserService
var auth helper.Auth

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
	auth = helper.SetupAuth("enhance-notes-2024")
	userService =service.NewUserService(mockRepository,auth)
	exitCode:=m.Run()
	os.Exit(exitCode)
}


func Test_ShouldCreateUserAndPassToken(t *testing.T){
	t.Run("ShouldCreateUserAndPassToken", func(t *testing.T) {
		initialUser := dto.UserSignUp{UserLogin: dto.UserLogin{Email:"test3@test.com", Password:"testtest123123"}} 
		token, _:=userService.SignUp(initialUser)
		bearerToken := "Bearer " + token
		user, err := auth.VerifyToken(bearerToken)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, initialUser.Email)	
	})
}
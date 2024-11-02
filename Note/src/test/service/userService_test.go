package service

/* import (
	"enhanced-notes/config"
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
		Password: "$2a$10$WVgDGLFlqvdjpRxSjrwv5u532KJp130adGrX0hOaKApe2VTV0M0P.",
		},
		{
		ID: 2,
		Email: "test1@test.com",
		Password: "testing123app",
		},
	}
	mockRepository := NewMockUserRepository(initialUsers)
	auth = helper.SetupAuth("enhance-notes-2024")
	cfg , _:= config.SetupEnv()
	userService =service.NewUserService(mockRepository,auth,cfg)
	exitCode:=m.Run()
	os.Exit(exitCode)
}


func Test_ShouldSignUpAndPassToken(t *testing.T){
	t.Run("ShouldCreateUserAndPassToken", func(t *testing.T) {
		initialUser := dto.UserSignUp{UserLogin: dto.UserLogin{Email:"test3@test.com", Password:"testtest123123"}}
		token, _:=userService.SignUp(initialUser)
		bearerToken := "Bearer " + token
		user, err := auth.VerifyToken(bearerToken)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, initialUser.Email)
	})
}

func Test_ShouldLoginAndPassToken(t *testing.T){
	t.Run("ShouldLoginAndPassToken", func(t *testing.T) {
		token, _:=userService.Login(dto.UserLogin{Email: "test@test.com", Password: "testing123app"} )
		bearerToken := "Bearer " + token
		user, err := auth.VerifyToken(bearerToken)
		assert.NoError(t, err)
		assert.Equal(t, "test@test.com", user.Email)
	})
}

func Test_ShouldFindUserByMail(t *testing.T){
	t.Run("ShouldFindUserByMail", func(t *testing.T) {
		user, err:=userService.FindUserByEmail("test@test.com")
		assert.NoError(t, err)
		assert.Equal(t, "test@test.com", user.Email)
	})
} */
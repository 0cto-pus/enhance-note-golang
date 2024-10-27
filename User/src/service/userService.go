package service

import (
	"enhanced-notes/config"
	"enhanced-notes/src/domain"
	"enhanced-notes/src/dto"
	"enhanced-notes/src/helper"
	"enhanced-notes/src/repository"
	"errors"
)

type IUserService interface {
    Login(userInput dto.UserLogin)(string, error)
    SignUp(userInput dto.UserSignUp)(string, error)
    FindUserByEmail(email string)(*domain.User,error)


}

type UserService struct {
	userRepository repository.IUserRepository
    Auth helper.Auth
    Config config.AppConfig
}



func NewUserService(userRepository repository.IUserRepository) IUserService{
    return &UserService{
        userRepository: userRepository,
    }
}


func(userService *UserService) Login(userInput dto.UserLogin)(string, error){
    user, err := userService.FindUserByEmail(userInput.Email)
    if err != nil {
        return "",errors.New("invalid credentials")
    }
    return userService.Auth.GenerateToken(user.ID, user.Email)
}

func(userService *UserService) SignUp(userInput dto.UserSignUp)(string, error){
    hashedPassword, err := userService.Auth.CreateHashedPassword(userInput.Password)
    if err != nil{
        return "",err
    }

    user, err := userService.userRepository.CreateUser(domain.User{Email:userInput.Email,Password: hashedPassword})

    if err != nil{
        return "",err
    }
    return userService.Auth.GenerateToken(user.ID, user.Email)
}

func (userService *UserService) FindUserByEmail(email string)(*domain.User,error){
    user,err := userService.userRepository.GetUserByEmail(email)
    
    return &user, err
}
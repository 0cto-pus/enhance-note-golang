package repository

import (
	"enhanced-notes/src/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(userId uint64) (domain.User, error)
}

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository{
	return &UserRepository{
		db:db,
	}
}

func (userRepository *UserRepository) CreateUser(user domain.User) (domain.User, error){
	err := userRepository.db.Create(&user).Error
	if err != nil{
		log.Printf("create user erorr: %v", err)
		return domain.User{}, errors.New("failed to create user")
	} 

	return user, nil
}

func (userRepository *UserRepository) FindUser(userId uint64) (domain.User,error){
	return domain.User{}, nil
}
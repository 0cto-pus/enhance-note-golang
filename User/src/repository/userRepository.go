package repository

import (
	"enhanced-notes/src/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user domain.User)
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

func (userRepository *UserRepository) CreateUser(user domain.User){
	
}

func (userRepository *UserRepository) FindUser(userId uint64) (domain.User,error){
	return domain.User{}, nil
}
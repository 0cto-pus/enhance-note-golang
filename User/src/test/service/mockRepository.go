package service

import (
	"enhanced-notes/src/domain"
	"enhanced-notes/src/repository"
	"fmt"
)


type MockUserRepository struct{
	users  []domain.User
}

func NewMockUserRepository(mockUsers []domain.User) repository.IUserRepository{
	return &MockUserRepository{
		users: mockUsers,
	}
}


func (userMockRepository *MockUserRepository) CreateUser(user domain.User) (domain.User, error){
	userMockRepository.users = append(userMockRepository.users, domain.User{
		ID:       uint64(len(userMockRepository.users)) + 1,
		Email:     user.Email,
		Password:    user.Password,
	})

	return  user,nil
}

func (userMockRepository *MockUserRepository) FindUserById(userId uint64) (domain.User,error){
	for _, user := range userMockRepository.users {
        if user.ID == userId {
            return user, nil
        }
    }
    return domain.User{}, fmt.Errorf("user with ID %v not found", userId)
}

func (userMockRepository *MockUserRepository) GetUserByEmail(email string) (domain.User, error) {
	for _, user := range userMockRepository.users {
        if user.Email == email {
            return user, nil
        }
    }
    return domain.User{}, fmt.Errorf("user with email %v not found", email)
}

package infra

import (
	"context"
	"enhanced-notes/src/domain"
	"enhanced-notes/src/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var userRepository repository.IUserRepository
var db *gorm.DB
var ctx context.Context

func TestMain(m *testing.M){

	var dsn string = "host=127.0.0.1 user=postgres password=root dbname=postgres port=6432 sslmode=disable"
	gormOpen, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	db = gormOpen;
	if err != nil {
		panic( err)
	}


	userRepository = repository.NewUserRepository(db)
	exitCode := m.Run()
	os.Exit(exitCode)

}
func setup(ctx context.Context, db *gorm.DB) {
	TestDataInitialize(ctx, db)
}
func clear(ctx context.Context, db *gorm.DB) {
	TruncateTestData(ctx, db)
}

func TestCreateUser(t *testing.T){

	 newUser := domain.User{
		Email: "test@test.com",
		Password: "hash-pass-mock",
	 }

	 t.Run("CreateUser", func(t *testing.T) {
		userRepository.CreateUser(newUser)
		addedUser, _ := userRepository.FindUserById(1)
		
		assert.Equal(t, domain.User{
			ID: 1,
			Email: "test@test.com",
			Password: "hash-pass-mock",
			CreatedAt: addedUser.CreatedAt,
			UpdatedAt: addedUser.UpdatedAt,}, addedUser)
	})
	clear(ctx, db)

}
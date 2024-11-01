package infra

import (
	"context"
	"enhanced-notes/src/domain"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func TestDataInitialize(ctx context.Context, db *gorm.DB) {
	testData := domain.User{
		Email: "test@test.com",
		Password: "hash-pass-mock",
	 }
	 err := db.Create(&testData).Error
	
	if err != nil {
		log.Error("Error inserting products: ", testData)
		return
	}
	log.Info(fmt.Sprintf("Products data created with %v rows", testData))
}
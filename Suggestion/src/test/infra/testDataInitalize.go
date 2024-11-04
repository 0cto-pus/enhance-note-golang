package infra

import (
	"context"
	"enhance-notes-suggestion/src/domain"

	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func TestDataInitialize(ctx context.Context, db *gorm.DB) {
	testData := domain.Suggestion{
		ID: 1,
		UserID: 1,
		NoteID: 1,
		Suggestion: "Deneme deneme", 
	 }
	 err := db.Create(&testData).Error

	if err != nil {
		log.Error("Error inserting products: ", testData)
		return
	}
	log.Info(fmt.Sprintf("Products data created with %v rows", testData))
} 
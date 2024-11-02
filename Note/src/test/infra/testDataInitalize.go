package infra

import (
	"context"
	"enhance-notes-note-service/src/domain"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func TestDataInitialize(ctx context.Context, db *gorm.DB) {
	testData := domain.Note{
		UserID: 1,
		Content: "Deneme Note",
	 }
	 err := db.Create(&testData).Error
	
	if err != nil {
		log.Error("Error inserting products: ", testData)
		return
	}
	log.Info(fmt.Sprintf("Products data created with %v rows", testData))
}
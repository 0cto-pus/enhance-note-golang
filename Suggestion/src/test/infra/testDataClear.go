package infra

import (
	"context"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func TruncateTestData(ctx context.Context, db *gorm.DB) {
	// 'Products' tablosunun içeriğini temizlemek için doğrudan SQL komutu kullanılır.
	truncateErr := db.WithContext(ctx).Exec("TRUNCATE TABLE suggestions RESTART IDENTITY CASCADE").Error
	if truncateErr != nil {
		log.Errorf("Failed to truncate notes table: %v", truncateErr)
	} else {
		log.Info("Products table truncated successfully")
	}
} 
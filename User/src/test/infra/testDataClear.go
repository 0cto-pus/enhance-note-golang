package infra

import (
	"context"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func TruncateTestData(ctx context.Context, db *gorm.DB) {
	// 'Products' tablosunun içeriğini temizlemek için doğrudan SQL komutu kullanılır.
	truncateErr := db.WithContext(ctx).Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
	if truncateErr != nil {
		log.Errorf("Failed to truncate products table: %v", truncateErr)
	} else {
		log.Info("Products table truncated successfully")
	}
}
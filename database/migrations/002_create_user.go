package migrations

import (
	"gorm.io/gorm"
	"rest/models"
)

func UpCreateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&models.Users{})

}

func DownCreateUsers(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Users{})
}

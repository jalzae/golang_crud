package migrations

import (
	"gorm.io/gorm"
	"rest/models"
)

func UpCreateBarang(db *gorm.DB) error {
	return db.AutoMigrate(&models.Barang{})

}

func DownCreateBarang(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Barang{})
}

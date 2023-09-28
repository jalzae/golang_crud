package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Barang struct {
	IdBarang       uuid.UUID `gorm:"type:char(36);primaryKey"`
	NamaBarang     string    `form:"NamaBarang" json:"NamaBarang" xml:"NamaBarang" binding:"required" gorm:"column:Nama_Barang;type:varchar(100);not null;"`
	UsersCreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`

	table string `gorm:"-"`
}

func (p Barang) TableName() string {
	// double check here, make sure the table does exist!!
	if p.table != "" {
		return p.table
	}
	return "Barang" // default table name
}

func CreateBarang(db *gorm.DB, Barangs *Barang) (err error) {
	err = db.Create(Barangs).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBarang(db *gorm.DB, Barangs *[]Barang) (err error) {
	err = db.Find(Barangs).Error
	if err != nil {
		return err
	}
	return nil
}

// get user by id
func GetBarangbyid(db *gorm.DB, Barangs *Barang, IdBarang string) (err error) {
	err = db.Where("id_barang = ?", IdBarang).First(Barangs).Error
	if err != nil {
		return err
	}
	return nil
}

// update user
func UpdateBarang(db *gorm.DB, Barangs *Barang) (err error) {
	db.Save(Barangs)
	return nil
}

// delete user
func DeleteBarang(db *gorm.DB, Barangs *Barang, IdBarang string) (err error) {
	db.Where("id_barang = ?", IdBarang).Delete(Barangs)
	return nil
}

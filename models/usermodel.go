package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rest/helper"
	"time"
)

type Users struct {
	UsersID             uuid.UUID `gorm:"type:char(36);primaryKey"`
	UsersName           string    `form:"UsersName" json:"UsersName" xml:"UsersName"  binding:"required" gorm:"type:varchar(100);not null;"`
	UsersCode           string    `form:"UsersCode" json:"UsersCode" xml:"UsersCode"  binding:"required" gorm:"type:varchar(100);not null;"`
	UsersPassword       string    `form:"UsersPassword" json:"UsersPassword" xml:"UsersPassword"  binding:"required" gorm:"type:text;not null"`
	UsersEmail          string    `form:"UsersEmail" json:"UsersEmail" xml:"UsersEmail"  binding:"required" gorm:"type:varchar(100);not null"`
	UsersStatus         int       `gorm:"type:int(10);not null"`
	UsersCreatedAt      time.Time `gorm:"type:datetime;sql:DEFAULT:current_timestamp;not null"`
	UsersCreatedBy      string    `gorm:"type:varchar(100);not null"`
	UsersLastmodifiedAt time.Time `gorm:"type:datetime;not null"`
	UsersLastmodifiedBy string    `gorm:"type:varchar(100);not null"`
	UsersUsertypesid    int       `gorm:"type:int(10);not null"`
	Token               string    `gorm:"type:text"`
}

func LoginUser(db *gorm.DB, User *Users, username string, password string) int64 {
	var err int64
	db.Table("users").Where("users_code", username).Where("users_password", password).Count(&err)
	return err
}

func CheckUser(db *gorm.DB, User *Users, username string) int64 {
	var err int64
	db.Table("users").Where("users_name", username).Count(&err)
	return err
}

func CreateUser(db *gorm.DB, User *Users) (err error) {
	if User.UsersID == uuid.Nil {
		User.UsersID = helper.GetUUID()
	}

	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB, User *[]Users) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get user by id
func GetUser(db *gorm.DB, User *Users, usersId string) (err error) {
	err = db.Where("users_id = ?", usersId).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// update user
func UpdateUser(db *gorm.DB, User *Users) (err error) {
	db.Save(User)
	return nil
}

// delete user
func DeleteUser(db *gorm.DB, User *Users, usersId string) (err error) {
	db.Where("users_id = ?", usersId).Find(&User)
	db.Unscoped().Model(&User).Update("deleted_at", nil)
	db.Unscoped().Delete(&User)
	return nil
}

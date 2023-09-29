package migrations

import (
    "gorm.io/gorm"
		"rest/models"
)

func UpRemoveUsersProfilesID(db *gorm.DB) error {
    // Remove the 'users_profilesid' column from the 'users' table
    return db.Migrator().DropColumn(&models.Users{}, "users_profilesid")
}

func DownRemoveUsersProfilesID(db *gorm.DB) error {
    // Add the 'users_profilesid' column back to the 'users' table
    return db.Migrator().AddColumn(&models.Users{}, "users_profilesid")
}

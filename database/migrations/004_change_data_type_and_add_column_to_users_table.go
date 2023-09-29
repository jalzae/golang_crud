package migrations

import (
	"gorm.io/gorm"
)

func UpChangeDataTypeAndAddColumn(db *gorm.DB) error {
	// Change the data type of the 'UsersUsertypesid' column to integer(1)
	sql := "ALTER TABLE users MODIFY COLUMN users_usertypesid INT(1);"
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	// Add a new column 'token_web' of type string
	sql = "ALTER TABLE users ADD COLUMN token_web VARCHAR(255);"
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	return nil
}

func DownChangeDataTypeAndAddColumn(db *gorm.DB) error {
	// Revert changes: change the data type of 'UsersUsertypesid' and remove 'token_web'
	sql := "ALTER TABLE users MODIFY COLUMN users_usertypesid INT(10);"
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	sql = "ALTER TABLE users DROP COLUMN token_web;"
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	return nil
}

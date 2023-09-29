package system

import (
	"fmt"
	"gorm.io/gorm"
	"rest/config"
	migrate "rest/database/migrations"
	"rest/models"
)

var Migrations = []struct {
	Name string
	Up   func(*gorm.DB) error
	Down func(*gorm.DB) error
}{
	{"001_create_barang", migrate.UpCreateBarang, migrate.DownCreateBarang},
	{"002_create_user", migrate.UpCreateUsers, migrate.DownCreateUsers},
	{"003_remove_users_profilesid_on_users", migrate.UpRemoveUsersProfilesID, migrate.DownRemoveUsersProfilesID},
	{"004_change_data_type_and_add_column_to_users_table", migrate.UpChangeDataTypeAndAddColumn, migrate.DownChangeDataTypeAndAddColumn},
	// Add more migrations as needed
}

func Migrate() {
	db := config.InitDb()

	// Run the migrations
	if err := RunMigrations(db); err != nil {
		panic("Error running migrations: " + err.Error())
	}
}

func RunMigrations(db *gorm.DB) error {

	for _, migration := range Migrations {
		applied := models.Migration{}
		db.Where("name = ?", migration.Name).First(&applied)

		if applied.ID == 0 {
			// Migration has not been applied, apply it
			if err := migration.Up(db); err != nil {
				return err
			}
			db.Create(&models.Migration{Name: migration.Name, Applied: true})
		}
	}

	return nil
}

func RollbackMigration(db *gorm.DB, filename string) error {
	// Find the migration by filename
	var targetMigration *models.Migration
	db.Where("name = ?", filename).First(&targetMigration)

	if targetMigration.ID == 0 {
		return fmt.Errorf("migration file '%s' not found", filename)
	}

	// Get the corresponding migration function
	var migrationFunc func(*gorm.DB) error
	for _, migration := range Migrations {
		if migration.Name == filename {
			migrationFunc = migration.Down
			break
		}
	}

	if migrationFunc == nil {
		return fmt.Errorf("migration file '%s' not supported for rollback", filename)
	}

	// Execute the "down" function of the target migration
	if err := migrationFunc(db); err != nil {
		return err
	}

	// Delete the migration record from the migration tracking table
	if err := db.Delete(targetMigration).Error; err != nil {
		return err
	}

	return nil
}

package db

import (
	"band-manager-backend/internal/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// dropDB drops all database tables (used for development/testing).
func dropDB() {
	db.Migrator().DropTable(
		&model.Group{},
		&model.User{},
		&model.Subgroup{},
		&model.UserGroupRole{},
		&model.Announcement{},
		&model.Event{},
		&model.Track{},
		&model.Notesheet{},
		"user_group",
		"subgroup_user",
		"notesheet_subgroup",
		"announcement_subgroup",
		"event_tracks",
		"event_users",
	)
	fmt.Println("all tables dropped successfully")
}

// createDB creates all required database tables using GORM AutoMigrate.
func createDB() {
	err := db.AutoMigrate(
		&model.Group{},
		&model.User{},
		&model.Subgroup{},
		&model.UserGroupRole{},
		&model.Announcement{},
		&model.Event{},
		&model.Track{},
		&model.Notesheet{},
		&model.GoogleToken{},
		&model.GoogleCalendarEvent{},
	)
	if err != nil {
		log.Fatal("migrations failed: ", err)
	}
	fmt.Println("migrations completed successfully")
}

// InitDB initializes the database connection and performs migrations.
// It uses environment variables for connection details and creates
// required tables using GORM AutoMigrate.
func InitDB() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("backend_manager_db connection failed")
	}

	db = database

	// Najpierw usuwamy wszystkie tabele
	//dropDB()
	// Potem tworzymy je od nowa
	createDB()

	fmt.Println("backend_manager_db connection successful")
}

// GetDB returns the global database connection instance.
func GetDB() *gorm.DB {
	return db
}

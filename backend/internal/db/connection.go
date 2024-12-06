package db

import (
	"band-manager-backend/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func createDB() {
	err := db.AutoMigrate(
		&models.Group{},
		&models.User{},
		&models.Subgroup{},
		&models.Announcement{},
		&models.Event{},
		&models.Track{},
		&models.Notesheet{},
	)
	if err != nil {
		log.Fatal("migrations failed: ", err)
	}
	fmt.Println("migrations completed successfully")
}

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

	if !db.Migrator().HasTable(&models.User{}) {
		createDB()
	}

	fmt.Println("backend_manager_db connection successful")
}

func GetDB() *gorm.DB {
	return db
}

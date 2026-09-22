package database

import (
	"fmt"
	"log"
	"todo/internal/config"
	"todo/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database successfully")

	if err := db.AutoMigrate(&models.Todos{}); err != nil {
		log.Fatal("Failed to migrate to database: ", err)
	}

	log.Println("Migrated to database")

	return db
}

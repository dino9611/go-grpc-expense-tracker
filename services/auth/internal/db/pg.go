package db

import (
	"fmt"
	"grpc-finance-app/services/auth/internal/config"
	"grpc-finance-app/services/auth/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPgDb(cfg config.Config) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("connection db error")
	}
	return db
}

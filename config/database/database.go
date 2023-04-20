package database

import (
	"context"
	"os"
	"time"

	"github.com/muhangga/internal/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Photo{}, &entity.Comment{}, &entity.SocialMedia{})

	pool, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get DB pool")
	}

	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(100)
	pool.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to ping DB")
	}

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get DB pool")
	}

	if err := sqlDB.Close(); err != nil {
		log.Error().Err(err).Msg("Failed to close DB")
	}
}

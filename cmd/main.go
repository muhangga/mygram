package main

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/muhangga/config"
	"github.com/muhangga/config/database"
	"github.com/muhangga/internal/router"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Error().Err(err).Msg("Failed to load env")
	}

	db := database.InitDB()
	defer database.CloseDB(db)

	cfg := config.NewConfig()
	server := router.NewServer(cfg)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		server.RunServer()
	}()
	wg.Wait()
}

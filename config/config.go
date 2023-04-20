package config

import (
	"github.com/muhangga/config/database"
	"gorm.io/gorm"
)

type Config interface {
	Database() *gorm.DB
}

type config struct{}

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return database.InitDB()
}

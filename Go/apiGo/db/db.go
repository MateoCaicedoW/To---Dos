package db

import (
	"github.com/mateo/apiGo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:1234@localhost:5432/api"
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Player{})
	database.AutoMigrate(&models.Team{})

	database.AutoMigrate(&models.PlayerTeam{})
	return database
}

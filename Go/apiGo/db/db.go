package db

import (
	"github.com/mateo/apiGo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:1234@localhost:5432/api"
	dbb, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbb.AutoMigrate(&models.Player{})
	dbb.AutoMigrate(&models.Team{})

	dbb.AutoMigrate(&models.PlayerTeam{})
	return dbb
}

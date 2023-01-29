package database

import (
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/models"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func DbMigrate(dbParam *gorm.DB) {
	DbConnection = dbParam

	DbConnection.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Store{},
		&models.Category{},
		&models.Product{},
		&models.ProductLog{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.UserCredentials{},
	)
}

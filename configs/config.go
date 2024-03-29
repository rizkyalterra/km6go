package configs

import (
	"batch6/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123ABC4d.@tcp(localhost:3306)/km6go?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	AutoMigration()
}

func AutoMigration() {
	DB.AutoMigrate(&models.News{}, &models.User{})
}

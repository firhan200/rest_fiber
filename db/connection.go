package db

import (
	"fmt"
	"os"

	"github.com/firhan200/rest_fiber/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_POST")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/rest_fiber?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect to database")
	}

	return db, nil
}

func Migrate() {
	db, _ := GetConnection()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})
}

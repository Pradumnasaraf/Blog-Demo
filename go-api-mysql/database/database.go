package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDBClient() {

	err := godotenv.Load()
	if err != nil {
		log.Print(".env file not found, trying to get from environment")
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DB)
	fmt.Println(err)
	DB.AutoMigrate(&Schedule{})
}

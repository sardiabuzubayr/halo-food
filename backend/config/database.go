package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := GetEnv("HOST")
	port := GetEnv("PORT")
	dbname := GetEnv("DB")
	username := GetEnv("DB_USERNAME")
	password := GetEnv("DB_PASSWORD")

	// dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}
	DB.AutoMigrate()
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load .env file")
	}
	return os.Getenv(key)
}

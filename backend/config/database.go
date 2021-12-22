package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	host := GetEnv("HOST")
	port := GetEnv("PORT")
	dbname := GetEnv("DB")
	username := GetEnv("DB_USERNAME")
	password := GetEnv("DB_PASSWORD")

	// dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}
	db.AutoMigrate()

	return db
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load .env file")
	}
	return os.Getenv(key)
}

package database

import (
	"fmt"
	"log"
	"report-generator/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	var (
		DbUsername = config.DbUsername
		DbPassword = config.DbPassword
		DbName     = config.DbName
		DbHost     = config.DbHost
		DbPort     = config.DbPort
	)

	dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database error: ")
		fmt.Println(err)
		return nil
	}

	return db
}

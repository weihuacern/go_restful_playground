package db

import (
	"../models"
	"../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
)

var db *gorm.DB
var err error

func ModelDBMigrate(thismodel interface{}) {
	if !db.HasTable(thismodel) {
		err := db.CreateTable(thismodel)
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(thismodel)
}

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	user := utils.GetEnv("PG_USER", "helios")
	password := utils.GetEnv("PG_PASSWORD", "helios")
	host := utils.GetEnv("PG_HOST", "localhost")
	port := utils.GetEnv("PG_PORT", "5432")
	database := utils.GetEnv("PG_DB", "bedb")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	ModelDBMigrate(&models.Task{})
	ModelDBMigrate(&models.User{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

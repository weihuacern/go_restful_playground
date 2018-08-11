package db

import (
	"../models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

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
	user := getEnv("PG_USER", "helios")
	password := getEnv("PG_PASSWORD", "helios")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "bedb")

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
	ModelDBMigrate(&models.AppPortal{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

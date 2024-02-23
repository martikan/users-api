package main

import (
	"fmt"
	"github.com/martikan/users-api/api"
	"github.com/martikan/users-api/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	log.Println("Firing up the API...")
}

func init() {
	// Read env file
	initConfig()

	// Init DB
	db := initDB()

	// Start API
	server, err := api.NewServer(db)
	if err != nil {
		log.Fatalln("Cannot create api: ", err)
	}

	port := viper.GetString("port")
	if err = server.Start(fmt.Sprintf("localhost:%s", port)); err != nil {
		log.Fatalln("Cannot start api: ", err)
	}
}

func initConfig() {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("env")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error during reading env file: ", err)
	}

	// Set defaults
	viper.SetDefault("port", "3000")
}

func initDB() *gorm.DB {
	dbUrl := viper.GetString("database.url")
	if dbUrl == "" {
		log.Fatalln("DATABASE_URL env variable must be provided")
	}
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect database")
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	// Migrate schema
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalln("Cannot migrate users table")
	}

	return db
}

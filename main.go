package main

import (
	"dekawaf/database"
	"dekawaf/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title Swagger for [Backend API Deka-waf]
// @version 1.0
// @description This is a document for API use in [Backend API Deka-waf]
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host      ai.cloudeka.id
// @BasePath  /core/dekawaf
// @schemes   https
func main() {
	// Environment variables are loaded from Kubernetes config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbEvent := os.Getenv("DB_EVENT")
	if dbEvent == "rollback_migrate" || dbEvent == "rollback" {
		database.RunRollback()
	}
	if dbEvent == "migrate_only" {
		database.RunMigration()
	}
	if dbEvent == "migrate" || dbEvent == "rollback_migrate" {
		database.RunMigration()
		database.RunSeeder()
	}

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}

}

package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.elastic.co/apm/module/apmgorm/v2"
	_ "go.elastic.co/apm/module/apmgorm/v2/dialects/postgres"
)

var (
	psqlConn *gorm.DB
	err      error
)

// initialize database
func init() {
	// // Environment variables are loaded from Kubernetes config
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	setupPostsqlConn()
}

func setupPostsqlConn() {

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("POSTGRE_HOST"), os.Getenv("POSTGRE_PORT"), os.Getenv("POSTGRE_USER"), os.Getenv("POSTGRE_DBNAME"), os.Getenv("POSTGRE_SSLMODE"), os.Getenv("POSTGRE_PASS"))
	psqlConn, err = apmgorm.Open("postgres", dns)

	if err != nil {
		log.Fatalf("Failed to create a connection to database error:%s", err)
	}

	err = psqlConn.DB().Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %s", err)
	}

	psqlConn.LogMode(true)

}

// PostsqlConn return mysql connection from gorm ORM
func PostsqlConn() *gorm.DB {
	return psqlConn
}

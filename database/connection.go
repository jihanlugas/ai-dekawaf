package database

import (
	"log"
	"net"
	"net/url"
	"os"

	"go.elastic.co/apm/module/apmsql/v2/pgxv5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var psqlConn *gorm.DB

// initialize database
func init() {
	// Environment variables are loaded from Kubernetes config
	// setupPostsqlConn()
}

func setupPostsqlConn() {
	dsnURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(os.Getenv("POSTGRE_USER"), os.Getenv("POSTGRE_PASS")),
		Host:   net.JoinHostPort(os.Getenv("POSTGRE_HOST"), os.Getenv("POSTGRE_PORT")),
		Path:   "/" + os.Getenv("POSTGRE_DBNAME"),
	}
	query := dsnURL.Query()
	query.Set("sslmode", os.Getenv("POSTGRE_SSLMODE"))
	dsnURL.RawQuery = query.Encode()
	dsn := dsnURL.String()

	var err error

	psqlConn, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:        dsn,
			DriverName: apmpgxv5.DriverName,
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		log.Fatalf("Failed to create a connection to database: %v", err)
	}

	db, err := psqlConn.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

// PostsqlConn returns PostgreSQL connection from GORM ORM.
func PostsqlConn() *gorm.DB {
	return psqlConn
}

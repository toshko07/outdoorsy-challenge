package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/toshko07/outdoorsy-challenge/internal/configs"
)

func Connect(config configs.DB) *sql.DB {
	log.Info("initializing database connection ...")
	database := newDBConnection(config)
	if err := database.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	return database
}

func LoadTestData(database *sql.DB, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err := database.Exec(request)
		if err != nil {
			return err
		}
	}
	return nil
}

func newDBConnection(config configs.DB) *sql.DB {
	db, err := sql.Open("postgres", buildDBConnectionString(config))
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	log.Infof("successfully connected to database: %s", config.Name)
	return db
}

func buildDBConnectionString(config configs.DB) string {
	host := config.Host
	encodedPassword := url.QueryEscape(config.Password)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Username, encodedPassword,
		host, config.Port, config.Name,
		"disable")
}

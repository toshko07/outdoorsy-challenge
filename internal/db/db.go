package db

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
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

func SetupTestDb() (*sql.DB, func()) {
	// Create the Postgres TestContainer
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mdillon/postgis:11",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "testingwithrentals",
			"POSTGRES_PASSWORD": "postgres",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		panic(err)
	}

	p, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		panic(err)
	}

	testDB, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		"postgres", "postgres",
		"localhost", p.Port(), "testingwithrentals",
		"disable"))
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	// Wait for the database to be ready
	time.Sleep(5 * time.Second)

	if err := testDB.Ping(); err != nil {
		panic(err)
	}

	return testDB, func() {
		if err := postgresC.Terminate(ctx); err != nil {
			panic(err)
		}
	}
}

func SetupTestData(database *sql.DB) {
	_, err := database.Exec(`
	CREATE OR REPLACE FUNCTION truncate_tables(username IN VARCHAR) RETURNS void AS $$
	DECLARE
		statements CURSOR FOR
			SELECT tablename FROM pg_tables
			WHERE tableowner = username AND schemaname = 'public';
	BEGIN
		FOR stmt IN statements LOOP
			EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
		END LOOP;
	END;
	$$ LANGUAGE plpgsql;
	SELECT truncate_tables('postgres');
	`)
	if err != nil {
		panic(err)
	}

	err = loadTestData(database, "../db/test_data/sql-init.sql")
	if err != nil {
		panic(err)
	}
}

func loadTestData(database *sql.DB, path string) error {
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

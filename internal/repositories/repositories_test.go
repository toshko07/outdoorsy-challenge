package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/toshko07/outdoorsy-challenge/internal/db"
)

var database *sql.DB

func TestMain(m *testing.M) {
	var shutdown func()
	database, shutdown = setupTestDb()
	exitCode := m.Run()
	shutdown()
	os.Exit(exitCode)
}

func setupTestDb() (*sql.DB, func()) {
	// Create the Postgres TestContainer
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mdillon/postgis:11",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "testingwithrentals",
			"POSTGRES_PASSWORD": "postgres",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
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

	if err := testDB.Ping(); err != nil {
		panic(err)
	}

	err = db.LoadTestData(testDB, "../db/test_data/sql-init.sql")
	if err != nil {
		panic(err)
	}

	return testDB, func() {
		if err := postgresC.Terminate(ctx); err != nil {
			panic(err)
		}
	}
}

func truncateDb() {
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
}

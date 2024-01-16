package repositories

import (
	"database/sql"
	"os"
	"testing"

	"github.com/toshko07/outdoorsy-challenge/internal/db"
)

var database *sql.DB

func TestMain(m *testing.M) {
	var shutdown func()
	database, shutdown = db.SetupTestDb()
	db.SetupTestData(database)
	exitCode := m.Run()
	shutdown()
	os.Exit(exitCode)
}

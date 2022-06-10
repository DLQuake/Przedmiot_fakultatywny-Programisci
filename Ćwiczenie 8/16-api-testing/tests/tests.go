package tests

import (
	"testing"
	"os"
	"time"
	"github.com/jmoiron/sqlx"
	"sql"
)

func NewUnit(t *testing.T) (*sqlx.DB, func()) {
	t.Helper()
	databasetest.StartContainer(t)
	db, err := sqlx.Open(os.Getenv("DB_DRIVER"), "root:1234@tcp(localhost:33060)/unit_test?parseTime=true")
	if err != nil {
		t.Fatalf("opening database connection: %v", err)
	}
	dbo, err := sql.Open(os.Getenv("DB_DRIVER"), "root:1234@tcp(localhost:33060)/unit_test?parseTime=true")
	if err != nil {
		t.Fatalf("opening database connection: %v", err)
	}
	t.Log("waiting for database to be ready")
	// Wait for the database to be ready. Wait 100ms longer between each attempt.
	// Do not try more than 20 times.
	var pingError error
	maxAttempts := 20
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 1000 * time.Millisecond)
	}
	if pingError != nil {
		databasetest.StopContainer(t)
		t.Fatalf("waiting for database to be ready: %v", pingError)
	}
	if err := schema.Migrate(dbo); err != nil {
		dbo.Close()
		databasetest.StopContainer(t)
		t.Fatalf("migrating: %s", err)
	}
	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		db.Close()
		databasetest.StopContainer(t)
	}
	return db, teardown
}
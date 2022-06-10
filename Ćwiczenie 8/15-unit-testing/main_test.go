package main

import (
	"testing"
	"os"
	"time"
	"github.com/jmoiron/sqlx"
)

func TestMainA(t *testing.T) {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	db, teardown := tests.NewUnit(t)
	defer teardown()

	if err := schema.Seed(db); err != nil {
		t.Fatal(err)
	}

	user := userUnitTest.User{DB: db}
	t.Run("UserList", user.List)
	t.Run("UserCreate", user.Crud)
}

package databasetest

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func StartContainer(t *testing.T) {
	t.Helper()
	cmd := exec.Command("docker", "run", "-d", "--name", "unit_test", "--publish", "33060:3306", "--env", "MYSQL_ROOT_PASSWORD=1234", "--env", "MYSQL_DATABASE=unit_test", "mysql:5.7.20")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("could not start docker : %v", err)
	}
}

// StopContainer stops and removes the specified container.
func StopContainer(t *testing.T) {
	t.Helper()
	if err := exec.Command("docker", "container", "rm", "-f", "unit_test").Run(); err != nil {
		t.Fatalf("could not stop mysql container: %v", err)
	}
}
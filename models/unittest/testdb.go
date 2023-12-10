package unittest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Tiburso/GoManager/models/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func fatalTestError(fmtStr string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, fmtStr, args...)
	os.Exit(1)
}

func MainTest(m *testing.M) {
	searchDir, _ := os.Getwd()
	for searchDir != "" {
		if _, err := os.Stat(filepath.Join(searchDir, "go.mod")); err == nil {
			break // The "go.mod" should be the one for Gitea repository
		}
		if dir := filepath.Dir(searchDir); dir == searchDir {
			searchDir = "" // reaches the root of filesystem
		} else {
			searchDir = dir
		}
	}
	if searchDir == "" {
		panic("The tests should run in a Gitea repository, there should be a 'go.mod' in the root")
	}

	if err := CreateTestEngine(); err != nil {
		fatalTestError("Error creating test engine: %v\n", err)
	}

	exitStatus := m.Run()

	//TODO: Check how to do custom teardown

	os.Exit(exitStatus)
}

// Create Test Engine creates a new sqlite3 in-memory database for testing
// using gorm.Open("sqlite3", ":memory:")
func CreateTestEngine() error {
	var err error
	db.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		return err
	}

	// Migrate the schema
	return db.AutoMigrate()
}

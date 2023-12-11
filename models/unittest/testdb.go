package unittest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Tiburso/GoManager/models/db"
	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var fixtures *testfixtures.Loader

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

	dialect := db.DB.Dialector.Name()

	if dialect != "sqlite" {
		return fmt.Errorf("dialect is not sqlite: %s", dialect)
	}

	coreDB, err := db.DB.DB()

	if err != nil {
		return err
	}

	// Load fixtures
	fixtures, err = testfixtures.New(
		testfixtures.Database(coreDB),
		testfixtures.Dialect(dialect),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Directory("../../models/fixtures"),
	)

	if err != nil {
		return err
	}

	//TODO: Check if I need to migrate the schema
	return db.AutoMigrate()
}

// LoadFixtures loads the fixtures from the fixtures directory
func LoadFixtures() error {
	if fixtures == nil {
		return fmt.Errorf("fixtures not loaded")
	}

	return fixtures.Load()
}

func PrepareTestDatabase() error {
	if err := LoadFixtures(); err != nil {
		return err
	}

	return nil
}

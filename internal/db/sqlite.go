package db

import (
	"sync"

	"github.com/jmoiron/sqlx"

	// Import module for SQLite DB
	_ "modernc.org/sqlite"

	"github.com/aceberg/ClickAHabit/internal/check"
	"github.com/aceberg/ClickAHabit/internal/models"
)

var mu sync.Mutex

func connect(path string) *sqlx.DB {
	dbx, err := sqlx.Connect("sqlite", path)
	check.IfError(err)

	return dbx
}

func exec(path string, sqlStatement string) {

	mu.Lock()
	dbx := connect(path)
	_, err := dbx.Exec(sqlStatement)
	mu.Unlock()

	check.IfError(err)
}

// Select - select all from DB
func Select(path string, tabname string) (checks []models.Check) {

	mu.Lock()
	dbx := connect(path)

	sqlStatement := `SELECT * FROM '` + tabname + `' ORDER BY ID DESC;`

	err := dbx.Select(&checks, sqlStatement)
	mu.Unlock()

	check.IfError(err)

	return checks
}

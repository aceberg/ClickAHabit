package db

import (
	"fmt"

	"github.com/aceberg/ClickAHabit/internal/models"
)

// Create - create table if not exists
func Create(path string) {

	sqlStatement := `CREATE TABLE IF NOT EXISTS checks (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"NAME"		TEXT,
		"GR"		TEXT,
		"COLOR"		TEXT,
		"ICON"		TEXT,
		"PLACE"		INTEGER,
		"COUNT"		INTEGER,
		"LINK" 		TEXT
	);`
	exec(path, sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS weeks (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"NAME"		TEXT,
		"GR"		TEXT,
		"COLOR"		TEXT,
		"ICON"		TEXT,
		"PLACE"		INTEGER,
		"COUNT"		INTEGER,
		"LINK" 		TEXT
	);`
	exec(path, sqlStatement)
}

// Insert - insert into DB
func Insert(path string, tabname string, check models.Check) {

	sqlStatement := `INSERT INTO '%s' (DATE, NAME, GR, COLOR, ICON, PLACE, COUNT, LINK)
	VALUES ('%s','%s','%s','%s','%s','%d','%d','%s');`

	check.Group = quoteStr(check.Group)
	check.Name = quoteStr(check.Name)

	sqlStatement = fmt.Sprintf(sqlStatement, tabname, check.Date, check.Name, check.Group, check.Color, check.Icon, check.Place, check.Count, check.Link)

	exec(path, sqlStatement)
}

// Update - update DB
func Update(path string, tabname string, check models.Check, id int) {

	sqlStatement := `UPDATE '%s' SET DATE='%s', NAME='%s', GR='%s', COLOR='%s', ICON='%s', PLACE='%d', COUNT='%d', LINK='%s' WHERE ID='%d';`

	check.Group = quoteStr(check.Group)
	check.Name = quoteStr(check.Name)

	sqlStatement = fmt.Sprintf(sqlStatement, tabname, check.Date, check.Name, check.Group, check.Color, check.Icon, check.Place, check.Count, check.Link, id)

	exec(path, sqlStatement)
}

// Delete - delete
func Delete(path string, tabname string, id int) {

	sqlStatement := `DELETE FROM '%s' WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, tabname, id)

	exec(path, sqlStatement)
}

// Clear - delete all checks from table
func Clear(path string, tabname string) {
	sqlStatement := `DELETE FROM '%s';`
	sqlStatement = fmt.Sprintf(sqlStatement, tabname)
	exec(path, sqlStatement)
}

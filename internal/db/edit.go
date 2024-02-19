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
		"COUNT"		INTEGER
	);`
	exec(path, sqlStatement)
}

// Insert - insert into DB
func Insert(path string, check models.Check) {

	sqlStatement := `INSERT INTO checks (DATE, NAME, GR, COLOR, ICON, PLACE, COUNT)
	VALUES ('%s','%s','%s','%s','%s','%d','%d');`

	check.Group = quoteStr(check.Group)
	check.Name = quoteStr(check.Name)

	sqlStatement = fmt.Sprintf(sqlStatement, check.Date, check.Name, check.Group, check.Color, check.Icon, check.Place, check.Count)

	exec(path, sqlStatement)
}

// Update - update DB
func Update(path string, check models.Check, id int) {

	sqlStatement := `UPDATE checks SET DATE='%s', NAME='%s', GR='%s', COLOR='%s', ICON='%s', PLACE='%d', COUNT='%d' WHERE ID='%d';`

	check.Group = quoteStr(check.Group)
	check.Name = quoteStr(check.Name)

	sqlStatement = fmt.Sprintf(sqlStatement, check.Date, check.Name, check.Group, check.Color, check.Icon, check.Place, check.Count, id)

	exec(path, sqlStatement)
}

// Delete - delete
func Delete(path string, id int) {

	sqlStatement := `DELETE FROM checks WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// Clear - delete all checks from table
func Clear(path string) {
	sqlStatement := `DELETE FROM checks;`
	exec(path, sqlStatement)
}

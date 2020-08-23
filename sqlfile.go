// Package sqlfile provides a way to execute sql file easily
//
// For more usage see https://github.com/tanimutomo/sqlfile
package sqlfile

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
)

// Exec execute SQL statements written int the specified sql file
// Note that you cannot use comment out in the sql file
func Exec(db *sql.DB, filepath string) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(file), "\n")
	line := strings.Join(lines, "")
	stmts := strings.Split(line, ";")
	stmts = stmts[:len(stmts)-1]

	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf(err.Error() + " : when executing > " + stmt)
		}
	}

	return nil
}

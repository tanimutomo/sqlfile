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
func Exec(
	db *sql.DB,
	filepath string,
) (
	res []sql.Result,
	err error,
) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return res, err
	}

	ls := strings.Split(string(f), "\n")

	var vls []string
	for _, l := range ls {
		comsep := strings.Split(l, "--")
		vls = append(vls, comsep[0])
	}

	l := strings.Join(vls, "")
	qs := strings.Split(l, ";")
	qs = qs[:len(qs)-1]

	var rs []sql.Result
	for _, q := range qs {
		r, err := db.Exec(q)
		if err != nil {
			return res, fmt.Errorf(err.Error() + " : when executing > " + q)
		}
		rs = append(rs, r)
	}

	return rs, nil
}

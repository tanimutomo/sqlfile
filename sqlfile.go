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

	var ncls []string
	for _, l := range ls {
		ncl := excludeComment(l)
		ncls = append(ncls, ncl)
	}

	l := strings.Join(ncls, "")
	qs := strings.Split(l, ";")
	qs = qs[:len(qs)-1]

	var rs []sql.Result
	for i, q := range qs {
		fmt.Println(i, q)
		r, err := db.Exec(q)
		if err != nil {
			return res, fmt.Errorf(err.Error() + " : when executing > " + q)
		}
		rs = append(rs, r)
	}

	return rs, nil
}

func excludeComment(line string) string {
	d := "\""
	s := "'"
	c := "--"

	var nc string
	ck := line
	mx := len(line) + 1

	for {
		if len(ck) == 0 {
			return nc
		}

		di := strings.Index(ck, d)
		si := strings.Index(ck, s)
		ci := strings.Index(ck, c)

		if di < 0 {
			di = mx
		}
		if si < 0 {
			si = mx
		}
		if ci < 0 {
			ci = mx
		}

		var ei int

		if di < si && di < ci {
			nc += ck[:di+1]
			ck = ck[di+1:]
			ei = strings.Index(ck, d)
		} else if si < di && si < ci {
			nc += ck[:si+1]
			ck = ck[si+1:]
			ei = strings.Index(ck, s)
		} else if ci < di && ci < si {
			return nc + ck[:ci]
		} else {
			return nc + ck
		}

		nc += ck[:ei+1]
		ck = ck[ei+1:]
	}
}

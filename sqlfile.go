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

// SqlFile represents a queries holder
type SqlFile struct {
	queries []string
}

// Exec execute SQL statements written int the specified sql file
func (s *SqlFile) Exec(db *sql.DB) (res []sql.Result, err error) {
	tx, err := db.Begin()
	if err != nil {
		return res, err
	}
	defer saveTx(tx, &err)

	var rs []sql.Result
	for _, q := range s.queries {
		r, err := tx.Exec(q)
		if err != nil {
			return res, fmt.Errorf(err.Error() + " : when executing > " + q)
		}
		rs = append(rs, r)
	}

	return rs, err
}

// Load load sql file from path, and return SqlFile pointer
func Load(path string) (*SqlFile, error) {
	ls, err := readFileByLine(path)
	if err != nil {
		return nil, err
	}

	var ncls []string
	for _, l := range ls {
		ncl := excludeComment(l)
		ncls = append(ncls, ncl)
	}

	l := strings.Join(ncls, "")
	qs := strings.Split(l, ";")
	qs = qs[:len(qs)-1]

	sqlfile := &SqlFile{
		queries: qs,
	}
	return sqlfile, nil
}

func readFileByLine(path string) (ls []string, err error) {
	f, err := ioutil.ReadFile((path))
	if err != nil {
		return ls, err
	}

	ls = strings.Split(string(f), "\n")
	return ls, nil
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

func saveTx(tx *sql.Tx, err *error) {
	if p := recover(); p != nil {
		tx.Rollback()
		panic(p)
	} else if *err != nil {
		tx.Rollback()
	} else {
		e := tx.Commit()
		err = &e
	}
}

package sqlfile

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestExec_Commit(t *testing.T) {
	t.Helper()
	db, mock := newMockDB(t)
	defer db.Close()

	tests := []struct {
		query  string
		lastID int64
		rows   int64
	}{
		{`DROP TABLE IF EXISTS users`, 0, 0},
		{`CREATE TABLE users (id BIGINT PRIMARY KEY AUTO_INCREMENT NOTNULL, name VARCHAR(255))`, 0, 0},
		{`INSERT INTO users (id, name) VALUES (1, 'user')`, 1, 1},
	}

	var qs []string
	for _, test := range tests {
		mock.ExpectExec(regexp.QuoteMeta(test.query)).
			WillReturnResult(sqlmock.NewResult(test.lastID, test.rows))
		qs = append(qs, test.query)
	}

	s := SqlFile{queries: qs}

	if _, err := s.Exec(db); err != nil {
		t.Errorf("test error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

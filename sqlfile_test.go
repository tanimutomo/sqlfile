package sqlfile

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
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

	mock.ExpectBegin()

	var qs []string
	for _, test := range tests {
		mock.ExpectExec(regexp.QuoteMeta(test.query)).
			WillReturnResult(sqlmock.NewResult(test.lastID, test.rows))
		qs = append(qs, test.query)
	}

	mock.ExpectCommit()

	s := SqlFile{queries: qs}

	if _, err := s.Exec(db); err != nil {
		t.Errorf("test error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestExec_Rollback(t *testing.T) {
	t.Helper()
	db, mock := newMockDB(t)
	defer db.Close()

	query := `INSERT INTO non_existing_table (id) values (1)`
	qs := []string{query}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnError(errors.New("Error 1146: Table 'tmp.non_existing_table' doesn't exist"))
	mock.ExpectRollback()

	s := SqlFile{queries: qs}

	s.Exec(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFile_SqlNotIncludeComments(t *testing.T) {
	t.Helper()

	exps, err := readFileByLine("./testdata/expected.sql")
	if err != nil {
		t.Fatalf(err.Error())
	}

	s := New()
	if err := s.File("./testdata/not_include_comments.sql"); err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, len(exps), len(s.queries))

	for i := 0; i < len(exps); i++ {
		assert.Equal(t, exps[i], s.queries[i])
	}
}

func TestFile_SqlIncludeComments(t *testing.T) {
	t.Helper()

	exps, err := readFileByLine("./testdata/expected.sql")
	if err != nil {
		t.Fatalf(err.Error())
	}

	s := New()
	if err := s.File("./testdata/include_comments.sql"); err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, len(exps), len(s.queries))

	for i := 0; i < len(exps); i++ {
		assert.Equal(t, exps[i], s.queries[i])
	}
}

func TestFile_NotFound(t *testing.T) {
	t.Helper()

	s := New()
	err := s.File("./testdata/non_exisiting.sql")

	assert.NotEqual(t, nil, err)
}

func TestFiles_Success(t *testing.T) {
	t.Helper()

	s := New()
	err := s.Files(
		"./testdata/include_comments.sql",
		"./testdata/not_include_comments.sql",
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, 2, len(s.files))
}

func TestFiles_NotFound(t *testing.T) {
	t.Helper()

	s := New()
	err := s.Files(
		"./testdata/non_exisiting.sql",
		"./testdata/non_exisiting.sql",
	)

	assert.NotEqual(t, nil, err)
}

func TestDirectory_Success(t *testing.T) {
	t.Helper()

	s := New()
	err := s.Directory("./testdata")
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, 3, len(s.files))
}

func TestDirectory_NotFound(t *testing.T) {
	t.Helper()

	s := New()
	err := s.Directory(
		"./non_exisiting",
	)

	assert.NotEqual(t, nil, err)
}

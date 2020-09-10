# sqlfile

![CI/CD](https://github.com/tanimutomo/sqlfile/workflows/gotest/badge.svg)
[![codecov](https://codecov.io/gh/tanimutomo/sqlfile/branch/master/graph/badge.svg)](https://codecov.io/gh/tanimutomo/sqlfile)
[![goreport](https://goreportcard.com/badge/github.com/tanimutomo/sqlfile)](https://goreportcard.com/report/github.com/tanimutomo/sqlfile)

A Golang library for treating SQL file.

sqlfile can execute multiple queries defined in .sql file with `database/sql`

## Installation
```
go get github.com/tanimutomo/sqlfile
```

## Usage
SQL) Prepare sql file.

Don't forget add `;` at last of each query.
```sql
-- example.sql

INSERT INTO users ( -- users table
  id, name, email, created_at, updated_at
) VALUES (
  1, 'user1', 'user1@example.com', now(), now() 
);

INSERT INTO articles ( -- articles table
  id, user_id, title, content, created_at, updated_at
) VALUES (
  1, 1, 'title1', "-- About -- \n I'm sqlfile.", now(), now() -- post1
), (
  2, 1, 'title2', '- About - \n I''m sqlfile.', now(), now() -- post2
);
```

Go) Load and Execute sql file.
```go
import (
  "database/sql"
  "github.com/tanimutomo/sqlfile"
)

// Get a database handler
db, err := sql.Open("DBMS", "CONNECTION")

// Initialize SqlFile
s := sqlfile.New()

// Load input file and store queries written in the file
err := s.File("example.sql")

// Load input files and store queries written in the files
err := s.Files("example.sql", "example2.sql")

// Load files in the input directory and store queries written in the files
err := s.Directory("./examples")

// Execute the stored queries
// transaction is used to execute queries in Exec()
res, err := s.Exec(db)
```

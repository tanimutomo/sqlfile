# sqlfile
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

// Load *.sql file
s, err := sqlfile.Load("example.sql")

// Execute queries in the loaded file
res, err := s.Exec(db)
```

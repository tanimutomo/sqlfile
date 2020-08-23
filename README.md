# sqlfile
**WIP**

A Golang library for using SQL file.

## Installation
```
go get github.com/tanimutomo/sqlfile
```

## Usage
Prepare sql file.
```
-- example.sql
INSERT INTO users ( -- users table
  id, name, email, created_at, updated_at
) VALUES (
  1, 'user1', 'user1@example.com', now(), now() 
);

INSERT INTO articles ( -- articles table
  id, user_id, title, content, created_at, updated_at
) VALUES (
  1, 1, 'title1', 'content1', now(), now() -- post1
), (
  2, 1, 'title2', 'content2', now(), now() -- post2
);
```

Execute sql file.
```
// Get a database handler
db, err := sql.Open("mysql", "connection settings")

res, err := sqlfile.Exec(db, "example.sql")
```

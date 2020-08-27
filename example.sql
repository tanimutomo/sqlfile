DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS articles;

CREATE TABLE users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  name VARCHAR(255),
  email VARCHAR(255),
  created_at DATETIME,
  updated_at DATETIME
);

CREATE TABLE articles (
  id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  user_id BIGINT,
  title VARCHAR(255),
  content TEXT,
  created_at DATETIME,
  updated_at DATETIME
);

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

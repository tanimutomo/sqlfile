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
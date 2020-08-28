-- drop tables --
DROP TABLE IF EXISTS users;-- users

-- create tables --
CREATE TABLE articles (-- articles
  id BIGINT,
  user_id BIGINT,
  title VARCHAR(255),
  content TEXT,
  created_at DATETIME,
  updated_at DATETIME
);

-- insert records --
INSERT INTO articles (-- articles
  id, user_id, title, content, created_at, updated_at
) VALUES (
  1, 1, 'Self Introduction', '-- sqlfile --\nI''m sqlfile.', now(), now()-- post 1
);
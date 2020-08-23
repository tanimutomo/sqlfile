INSERT INTO users ( -- users table
  id, name, email, created_at, updated_at
) VALUES (
  1, 'user1', 'user1@example.com', now(), now() 
);

INSERT INTO articles ( -- articles table
  id, title, content, created_at, updated_at
) VALUES (
  1, 'title1', 'content1', now(), now() -- post1
), (
  2, 'title2', 'content2', now(), now() -- post2
);
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

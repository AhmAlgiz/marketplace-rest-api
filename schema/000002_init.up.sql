ALTER TABLE users
RENAME COLUMN name TO username;

ALTER TABLE users 
ADD UNIQUE (username);
CREATE TABLE  users (
  id bigserial NOT NULL primary key,
  email varchar(255) NOT NULL unique,
  encrypted_password varchar(255) NOT NULL
)
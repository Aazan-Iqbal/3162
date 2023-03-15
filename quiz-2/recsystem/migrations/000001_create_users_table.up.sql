-- Filename: migrations/000001_create_users_table.up.sql

CREATE TABLE IF NOT EXISTS users (
  user_id bigserial PRIMARY KEY,
  email varchar(255) NOT NULL,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  dob date NOT NULL,
  address varchar(255) NOT NULL,
  phone_number varchar(255) NOT NULL,
  roles_id int NOT NULL REFERENCES "roles" ("roles_id"),
  password varchar(255) NOT NULL
);
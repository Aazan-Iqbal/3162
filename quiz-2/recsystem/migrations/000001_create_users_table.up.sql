-- Filename: migrations/000001_create_users_table.up.sql

CREATE TYPE address_type AS ENUM ('Belize', 'Santa Elena', 'San Ignacio', 'Belmopan', 'Dangriga');

CREATE TABLE IF NOT EXISTS users (
  user_id bigserial PRIMARY KEY,
  email varchar(255) NOT NULL,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  dob date NOT NULL,
  address address_type NOT NULL,
  phone_number varchar(255) NOT NULL,
  roles_id int REFERENCES "roles" ("roles_id"),
  password varchar(255) NOT NULL,
);
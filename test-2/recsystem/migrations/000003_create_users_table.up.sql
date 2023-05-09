-- Filename: migrations/000001_create_users_table.up.sql

CREATE TABLE IF NOT EXISTS "users" (
  "user_id" bigserial PRIMARY KEY,
  "email" varchar(255) NOT NULL,
  "first_name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "dob" text NOT NULL,
  "address" varchar(255) NOT NULL,
  "phone_number" text NOT NULL,
  "roles_id" int NOT NULL REFERENCES "roles" ("roles_id"),
  "password_hash" bytea NOT NULL,
  "created_at" TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()


);
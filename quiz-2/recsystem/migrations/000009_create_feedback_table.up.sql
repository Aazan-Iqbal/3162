CREATE TABLE IF NOT EXISTS "feedback" (
  "feedback_id" bigserial GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "users_id" bigserial REFERENCES "users" ("users_id"),
  "star_count" int NOT NULL,
  "message" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL
);
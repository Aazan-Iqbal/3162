CREATE TABLE IF NOT EXISTS "announcements" (
  "announcements_id" bigserial PRIMARY KEY,
  "subject" varchar(255) NOT NULL,
  "content" text NOT NULL
);

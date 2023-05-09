CREATE TABLE IF NOT EXISTS "equipment" (
  "equipment_id" bigserial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "image" bytea,
  "equipment_type_id" int NOT NULL REFERENCES "equipment_types" ("equipment_types_id"),
  "status" boolean NOT NULL,
  "availability" boolean NOT NULL 
);
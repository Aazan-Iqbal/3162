CREATE TABLE IF NOT EXISTS"equipment_usage_log" (
  "equipment_usage_log_id" bigserial PRIMARY KEY,
  "equipments_id" bigint NOT NULL REFERENCES "equipment" ("equipment_id"),
  "user_id" bigserial NOT NULL REFERENCES "users" ("user_id"), 
  "log_id" bigserial,
  "time_borrowed" TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  "returned_status" boolean NOT NULL
);
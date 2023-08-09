CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar,
  "access_level" int,
  "status" varchar,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "informations" (
  "id" SERIAL PRIMARY KEY,
  "access_level" int,
  "title" varchar,
  "content" varchar,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "authentications" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer,
  "value" varchar,
  "type" varchar,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "logs" ( -- for audit purposes
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "information_id" int,
  "ip_address" varchar,
  "action" varchar,
  "access_time" timestamp
);

ALTER TABLE "authentications" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "logs" ADD FOREIGN KEY ("information_id") REFERENCES "users" ("id");

CREATE INDEX "index_users_on_username" ON "users" ("username");
CREATE INDEX "index_informations_on_title" ON "informations" ("title");
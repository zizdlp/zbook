
CREATE TABLE "oauths" (
  "oauth_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "oauth_type" varchar(255) NOT NULL,
  "app_id" varchar NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "oauths"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "oauths" ("oauth_type", "app_id");  --- 添加索引，加快检索
CREATE UNIQUE INDEX ON "oauths" ("oauth_type", "user_id");  --- 添加索引，加快检索


CREATE TABLE "verifications" (
  "verification_id" BIGSERIAL PRIMARY KEY,
  "verification_url" VARCHAR(255) NOT NULL UNIQUE, -- 随机字符串 URL
  "verification_type" varchar(255) NOT NULL,
  "user_id" bigint NOT NULL,
  "is_used" boolean NOT NULL DEFAULT FALSE,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expired_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL '10 minutes')
);

ALTER TABLE "verifications"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

CREATE INDEX ON "verifications" ("user_id");
CREATE INDEX ON "verifications" ("verification_url");

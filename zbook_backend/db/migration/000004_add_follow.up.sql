CREATE TABLE "follows" (
  "follow_id" bigserial PRIMARY KEY,
  "follower_id" bigint NOT NULL,
  "following_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE "follows"
  ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("following_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

CREATE INDEX ON "follows" ("follower_id");
CREATE INDEX ON "follows" ("following_id");
CREATE UNIQUE INDEX ON "follows" ("follower_id","following_id");
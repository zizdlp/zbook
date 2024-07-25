CREATE TABLE "repo_relations" (
  "relation_id" bigserial PRIMARY KEY,
  "relation_type" varchar(255) NOT NULL,
  "user_id" bigint NOT NULL,
  "repo_id" bigint NOT NULL,
  "created_at"  timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "repo_relations"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("repo_id") REFERENCES "repos" ("repo_id") ON DELETE CASCADE;
CREATE UNIQUE INDEX ON "repo_relations" ("user_id","repo_id","relation_type");
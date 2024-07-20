CREATE TABLE "repos" (
  "repo_id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "git_protocol" varchar(255) NOT NULL,
  "git_host" varchar(255) NOT NULL,
  "git_username" varchar(255) NOT NULL,
  "git_repo" varchar(255) NOT NULL,
  "git_access_token" varchar(255) DEFAULT '',
  "repo_name" varchar(255) NOT NULL,
  "repo_description" text NOT NULL,
  "home_page" text NOT NULL,
  "sync_token" varchar(255) DEFAULT '',
  "visibility_level" varchar(255) NOT NULL,
  "commit_id" varchar(255) NOT NULL,
  "layout" text NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "fts_repo_name" TSVECTOR
);

ALTER TABLE "repos"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

-- 初始化 fts_repo_name
UPDATE "repos"
SET "fts_repo_name" = setweight(to_tsvector('english', "repo_name"), 'A');

-- 创建全文搜索索引
CREATE INDEX "repos_fts_repo_name_gin_index" ON "repos" USING gin ("fts_repo_name");

-- 在 repo_name 更新时触发更新 fts_repo_name
CREATE TRIGGER "trig_repos_update_fts_repo_name"
  BEFORE INSERT OR UPDATE OF "repo_name"
  ON "repos"
  FOR EACH ROW
  EXECUTE FUNCTION tsvector_update_trigger("fts_repo_name", 'pg_catalog.english', "repo_name");

CREATE UNIQUE INDEX ON "repos" ("user_id","repo_name");
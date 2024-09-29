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
  "sync_token" varchar(255) DEFAULT '',
  "visibility_level" varchar(255) NOT NULL,
  "commit_id" varchar(255) NOT NULL,
  "config" text NOT NULL DEFAULT '',
  "home" text NOT NULL DEFAULT '',
  "theme_sidebar" text NOT NULL CHECK (length(trim(theme_sidebar)) > 0),
  "theme_color" text NOT NULL CHECK (length(trim(theme_color)) > 0),
  "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "branch" varchar(255) NOT NULL DEFAULT '',
  "fts_repo_en" TSVECTOR,
  "fts_repo_zh" TSVECTOR
);

ALTER TABLE "repos"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "repos" ("user_id","repo_name");

-- 初始化 fts_repo_name_en
UPDATE "repos"
SET "fts_repo_en" = setweight(to_tsvector('english', "repo_name"), 'A') ||
                    setweight(to_tsvector('english', "repo_description"), 'B');

-- 初始化 fts_repo_name_zh
UPDATE "repos"
SET "fts_repo_zh" = setweight(to_tsvector('jiebacfg', "repo_name"), 'A') || 
                    setweight(to_tsvector('jiebacfg', "repo_description"), 'B');

-- 创建全文搜索索引
CREATE INDEX "repos_fts_repo_en_gin_index" ON "repos" USING gin ("fts_repo_en");
CREATE INDEX "repos_fts_repo_zh_gin_index" ON "repos" USING gin ("fts_repo_zh");

-- 在 repo_name,repo_description 更新时触发更新 fts_repo_en
CREATE TRIGGER "trig_repos_update_fts_repo_en"
  BEFORE INSERT OR UPDATE OF "repo_name","repo_description"
  ON "repos"
  FOR EACH ROW
  EXECUTE FUNCTION tsvector_update_trigger("fts_repo_en", 'pg_catalog.english', "repo_name","repo_description");

-- 在 repo_name,repo_description 更新时触发更新 fts_repo_zh
CREATE TRIGGER "trig_repos_update_fts_repo_zh"
  BEFORE INSERT OR UPDATE OF "repo_name","repo_description"
  ON "repos"
  FOR EACH ROW
  EXECUTE FUNCTION tsvector_update_trigger("fts_repo_zh", 'public.jiebacfg', "repo_name","repo_description");
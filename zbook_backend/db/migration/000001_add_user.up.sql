-- 创建 users 表
CREATE TABLE "users" (
  "user_id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(255) UNIQUE NOT NULL CHECK (char_length(username) >= 3),
  "email" VARCHAR(255) UNIQUE NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z]{2,}$'),
  "hashed_password" VARCHAR(255) NOT NULL,
  "blocked" BOOLEAN NOT NULL DEFAULT FALSE,
  "verified" BOOLEAN NOT NULL DEFAULT FALSE,
  "motto" TEXT NOT NULL DEFAULT 'Strive for progress, not perfection.',
  "user_role" VARCHAR(50) NOT NULL DEFAULT 'user',
  "onboarding" BOOLEAN NOT NULL DEFAULT TRUE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "unread_count" INT NOT NULL DEFAULT 0,
  "unread_count_updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "fts_username" TSVECTOR
);

-- 创建索引
CREATE INDEX "idx_users_name_email" ON "users" ("username", "email");

-- 创建扩展
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS btree_gin;
CREATE EXTENSION IF NOT EXISTS pg_jieba;

-- 初始化 fts_username
UPDATE "users"
SET "fts_username" = setweight(to_tsvector('english', "username"), 'A');

-- 创建全文搜索索引
CREATE INDEX "users_fts_username_gin_index" ON "users" USING gin ("fts_username");

-- 在 username 更新时触发更新 fts_username
CREATE TRIGGER "trig_users_update_fts_username"
  BEFORE INSERT OR UPDATE OF "username"
  ON "users"
  FOR EACH ROW
  EXECUTE FUNCTION tsvector_update_trigger("fts_username", 'pg_catalog.english', "username");

-- 创建触发器，通知未读消息数量的变化
CREATE OR REPLACE FUNCTION "notify_unread_count_change"() RETURNS TRIGGER AS $$
BEGIN
    PERFORM pg_notify('unread_count_change', json_build_object('username', NEW."username", 'unread_count', NEW."unread_count")::text);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 在 unread_count 更新时触发通知
CREATE TRIGGER "trig_users_unread_count_change"
    AFTER UPDATE OF "unread_count" ON "users"
    FOR EACH ROW
    EXECUTE FUNCTION "notify_unread_count_change"();


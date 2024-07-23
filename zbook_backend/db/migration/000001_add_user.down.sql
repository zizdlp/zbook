-- 删除表
DROP TABLE IF EXISTS "users";

-- 删除索引
DROP INDEX IF EXISTS "idx_users_name_email";
DROP INDEX IF EXISTS "users_fts_username_gin_index";

-- 删除扩展
DROP EXTENSION IF EXISTS pg_trgm;
DROP EXTENSION IF EXISTS btree_gin;
DROP EXTENSION IF EXISTS pg_jieba;

-- 删除触发器
DROP TRIGGER IF EXISTS "trig_users_update_fts_username" ON "users";
DROP TRIGGER IF EXISTS "trig_users_unread_count_change" ON "users";


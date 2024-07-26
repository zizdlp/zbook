DROP TABLE IF EXISTS "comments";

DROP INDEX IF EXISTS "markdowns_fts_comment_zh_gin_index";
DROP INDEX IF EXISTS "markdowns_fts_comment_en_gin_index";

-- 删除触发器
DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_comment_zh" ON "comments";
DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_comment_en" ON "comments";
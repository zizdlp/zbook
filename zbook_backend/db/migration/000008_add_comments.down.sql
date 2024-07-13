DROP TABLE IF EXISTS "comments";

DROP INDEX IF EXISTS "markdowns_fts_comment_content_gin_index";

-- 删除触发器
DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_comment_content" ON "comments";
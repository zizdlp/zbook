DROP TABLE IF EXISTS "comment_relations";
DROP TABLE IF EXISTS "comment_reports";

DROP INDEX IF EXISTS "markdowns_fts_report_content_gin_index";

-- 删除触发器
DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_report_content" ON "comment_reports";
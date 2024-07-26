DROP TABLE IF EXISTS "comment_relations";
DROP TABLE IF EXISTS "comment_reports";

DROP INDEX IF EXISTS "markdowns_fts_report_zh_gin_index";
DROP INDEX IF EXISTS "markdowns_fts_report_en_gin_index";

-- 删除触发器
DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_report_zh" ON "comment_reports";
DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_report_en" ON "comment_reports";
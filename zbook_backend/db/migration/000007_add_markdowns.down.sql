DROP TABLE IF EXISTS "markdowns";

DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_zh" ON "markdowns";
DROP TRIGGER IF EXISTS "trig_markdowns_insert_update_en" ON "markdowns";
DROP INDEX IF EXISTS "markdowns_fts_zh_gin_index";
DROP INDEX IF EXISTS "markdowns_fts_en_gin_index";

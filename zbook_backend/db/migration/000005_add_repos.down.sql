DROP TABLE IF EXISTS "repos";

DROP INDEX IF EXISTS "repos_fts_reop_name_gin_index";

-- 删除触发器
DROP TRIGGER IF EXISTS "trig_repos_update_fts_repo_zh" ON "repos";
DROP TRIGGER IF EXISTS "trig_repos_update_fts_repo_en" ON "repos";
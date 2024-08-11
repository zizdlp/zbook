CREATE TABLE "comments" (
  "comment_id" bigserial PRIMARY KEY,
  "repo_id" bigint NOT NULL,
  "markdown_id" bigint NOT NULL,
  "parent_id" bigint DEFAULT NULL,
  "root_id" bigint DEFAULT NULL,
  "user_id" bigint NOT NULL,
  "blocked" boolean NOT NULL DEFAULT 'false',
  "comment_content" text NOT NULL,
  "created_at"  timestamptz NOT NULL DEFAULT (now()),
  fts_comment_zh tsvector,
  fts_comment_en tsvector
);


-- Adding foreign key constraints
ALTER TABLE "comments" 
  ADD FOREIGN KEY ("markdown_id") REFERENCES "markdowns" ("markdown_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("parent_id") REFERENCES "comments" ("comment_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("repo_id") REFERENCES "repos" ("repo_id") ON DELETE CASCADE;

UPDATE "comments"
SET fts_comment_zh = setweight(to_tsvector('jiebacfg', "comment_content"), 'A');

UPDATE "comments"
SET fts_comment_en = setweight(to_tsvector('english', "comment_content"), 'A');

CREATE INDEX markdowns_fts_comment_zh_gin_index ON "comments" USING gin (fts_comment_zh);
CREATE INDEX markdowns_fts_comment_en_gin_index ON "comments" USING gin (fts_comment_en);


CREATE TRIGGER trig_markdowns_insert_update_comment_zh
  BEFORE INSERT OR UPDATE OF "comment_content"
  ON "comments"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_comment_zh, 'public.jiebacfg', "comment_content");


CREATE TRIGGER trig_markdowns_insert_update_comment_en
  BEFORE INSERT OR UPDATE OF "comment_content"
  ON "comments"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_comment_en, 'pg_catalog.english', "comment_content");

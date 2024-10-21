CREATE TABLE "markdowns" (
  "markdown_id" bigserial PRIMARY KEY,
  "relative_path" text NOT NULL,
  "user_id" bigint NOT NULL,
  "repo_id" bigint NOT NULL,
  "content" text NOT NULL,
  "updated_at"  timestamptz NOT NULL DEFAULT (now()),
  "created_at"  timestamptz NOT NULL DEFAULT (now()),
  fts_zh tsvector,
  fts_en tsvector
);

ALTER TABLE "markdowns"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("repo_id") REFERENCES "repos" ("repo_id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "markdowns" ("repo_id","relative_path");

UPDATE "markdowns"
SET fts_zh = setweight(to_tsvector('jiebacfg', "relative_path"), 'A') ||
          setweight(to_tsvector('jiebacfg', "content"), 'B');


UPDATE "markdowns"
SET fts_en = setweight(to_tsvector('english', "relative_path"), 'A') ||
          setweight(to_tsvector('english', "content"), 'B');


CREATE INDEX markdowns_fts_zh_gin_index ON "markdowns" USING gin (fts_zh);
CREATE INDEX markdowns_fts_en_gin_index ON "markdowns" USING gin (fts_en);

CREATE TRIGGER trig_markdowns_insert_update_zh
  BEFORE INSERT OR UPDATE OF "relative_path","content"
  ON "markdowns"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_zh, 'public.jiebacfg', "relative_path", "content");


CREATE TRIGGER trig_markdowns_insert_update_en
  BEFORE INSERT OR UPDATE OF "relative_path","content"
  ON "markdowns"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_en, 'pg_catalog.english', "relative_path", "content");

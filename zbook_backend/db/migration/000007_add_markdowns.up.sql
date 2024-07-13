CREATE TABLE "markdowns" (
  "markdown_id" bigserial UNIQUE NOT NULL,
  "relative_path" text NOT NULL,
  "user_id" bigint NOT NULL,
  "repo_id" bigint NOT NULL,
  "main_content" text NOT NULL,
  "table_content" text NOT NULL,
  "md5" varchar NOT NULL, -- 文件签名，用于判断markdown是否更新
  "version_key" varchar NOT NULL,
  "updated_at"  timestamptz NOT NULL DEFAULT (now()),
  "created_at"  timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("relative_path","repo_id")
);

ALTER TABLE "markdowns" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "markdowns" ADD FOREIGN KEY ("repo_id") REFERENCES "repos" ("repo_id");

CREATE UNIQUE INDEX ON "markdowns" ("repo_id","relative_path");


ALTER TABLE "markdowns" ADD COLUMN fts_zh tsvector;
ALTER TABLE "markdowns" ADD COLUMN fts_en tsvector;

UPDATE "markdowns"
SET fts_zh = setweight(to_tsvector('jiebacfg', "relative_path"), 'A') ||
          setweight(to_tsvector('jiebacfg', "main_content"), 'B');


UPDATE "markdowns"
SET fts_en = setweight(to_tsvector('english', "relative_path"), 'A') ||
          setweight(to_tsvector('english', "main_content"), 'B');


CREATE INDEX markdowns_fts_zh_gin_index ON "markdowns" USING gin (fts_zh);
CREATE INDEX markdowns_fts_en_gin_index ON "markdowns" USING gin (fts_en);

CREATE TRIGGER trig_markdowns_insert_update_zh
  BEFORE INSERT OR UPDATE OF "relative_path","main_content"
  ON "markdowns"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_zh, 'public.jiebacfg', "relative_path", "main_content");


CREATE TRIGGER trig_markdowns_insert_update_en
  BEFORE INSERT OR UPDATE OF "relative_path","main_content"
  ON "markdowns"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_en, 'pg_catalog.english', "relative_path", "main_content");

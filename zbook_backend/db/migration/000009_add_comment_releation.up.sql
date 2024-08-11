-- 1. comment_likes
CREATE TABLE "comment_relations" (
  "relation_id" bigserial PRIMARY KEY,
  "relation_type" varchar(255) NOT NULL,
  "user_id" bigint NOT NULL,
  "comment_id" bigint NOT NULL,
  "created_at"  timestamptz NOT NULL DEFAULT (now())
);

-- Adding foreign key constraints
ALTER TABLE "comment_relations"
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("comment_id") ON DELETE CASCADE;


CREATE UNIQUE INDEX ON "comment_relations" ("user_id","comment_id","relation_type");

-- 4. reports
CREATE TABLE "comment_reports" (
  "report_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "comment_id" bigint NOT NULL,
  "report_content" varchar NOT NULL DEFAULT '',
  "processed" boolean NOT NULL DEFAULT 'false',
  "created_at"  timestamptz NOT NULL DEFAULT (now()),
  fts_report_zh tsvector,
  fts_report_en tsvector
);

-- Adding foreign key constraints with ON DELETE CASCADE
ALTER TABLE "comment_reports" 
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE,
  ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("comment_id") ON DELETE CASCADE;

CREATE UNIQUE INDEX ON "comment_reports" ("user_id","comment_id");

UPDATE "comment_reports"
SET fts_report_zh = setweight(to_tsvector('jiebacfg', "report_content"), 'A');

UPDATE "comment_reports"
SET fts_report_en = setweight(to_tsvector('english', "report_content"), 'A');


CREATE INDEX markdowns_fts_report_zh_gin_index ON "comment_reports" USING gin (fts_report_zh);
CREATE INDEX markdowns_fts_report_en_gin_index ON "comment_reports" USING gin (fts_report_en);

CREATE TRIGGER trig_markdowns_insert_update_report_zh
  BEFORE INSERT OR UPDATE OF "report_content"
  ON "comment_reports"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_report_zh, 'public.jiebacfg', "report_content");

CREATE TRIGGER trig_markdowns_insert_update_report_en
  BEFORE INSERT OR UPDATE OF "report_content"
  ON "comment_reports"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_report_en, 'pg_catalog.english', "report_content");
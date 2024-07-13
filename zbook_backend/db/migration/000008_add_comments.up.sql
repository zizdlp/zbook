CREATE TABLE "comments" (
  "comment_id" bigserial PRIMARY KEY,
  "markdown_id" bigint NOT NULL,
  "parent_id" bigint DEFAULT NULL,
  "root_id" bigint DEFAULT NULL,
  "user_id" bigint NOT NULL,
  "blocked" boolean NOT NULL DEFAULT 'false',
  "comment_content" text NOT NULL,
  "created_at"  timestamptz NOT NULL DEFAULT (now()),
  "deleted" boolean NOT NULL DEFAULT 'false',
  fts_comment_content tsvector
);

ALTER TABLE "comments" ADD FOREIGN KEY ("markdown_id") REFERENCES "markdowns" ("markdown_id");
ALTER TABLE "comments" ADD FOREIGN KEY ("parent_id") REFERENCES "comments" ("comment_id");
ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

UPDATE "comments"
SET fts_comment_content = setweight(to_tsvector('jiebacfg', "comment_content"), 'A');

CREATE INDEX markdowns_fts_comment_content_gin_index ON "comments" USING gin (fts_comment_content);

CREATE TRIGGER trig_markdowns_insert_update_comment_content
  BEFORE INSERT OR UPDATE OF "comment_content"
  ON "comments"
  FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger(fts_comment_content, 'public.jiebacfg', "comment_content");

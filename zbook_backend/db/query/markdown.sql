-- name: CreateMarkdown :one
INSERT INTO markdowns (
  relative_path,
  user_id,
  repo_id,
  main_content,
  table_content
) VALUES (
  $1, $2, $3,$4,$5
) RETURNING *;

-- name: CreateMarkdownMulti :exec
INSERT INTO markdowns  (
  relative_path,
  user_id,
  repo_id,
  main_content,
  table_content
)
SELECT unnest(@relative_path::text[]) AS relative_path,
  unnest(@user_id::bigint[]) AS user_id,
  unnest(@repo_id::bigint[]) AS repo_id,
  unnest(@main_content::text[]) AS main_content,
  unnest(@table_content::text[]) AS table_content;


-- name: GetMarkdownContent :one
SELECT 
  markdowns.*
FROM markdowns
WHERE markdowns.relative_path = $1  and markdowns.repo_id = $2
LIMIT 1;

-- name: GetMarkdownByID :one
SELECT * FROM markdowns
WHERE markdown_id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: GetMarkdownRepoID :one
SELECT 
  markdowns.repo_id
FROM 
  markdowns
WHERE 
  markdown_id = $1;


-- name: UpdateMarkdownMulti :exec
UPDATE markdowns AS m
SET main_content=tmp.main_content,table_content=tmp.table_content,relative_path=new_relative_path,updated_at=now()
FROM (
  SELECT 
  unnest(@relative_path::text[]) AS relative_path,
  unnest(@new_relative_path::text[]) AS new_relative_path,
  unnest(@main_content::text[]) AS main_content,
  unnest(@table_content::text[]) AS table_content,
  unnest(@repo_id::bigint[]) AS repo_id
) AS tmp
WHERE m.relative_path = tmp.relative_path and m.repo_id=tmp.repo_id;

-- name: DeleteMarkdownMulti :exec
DELETE FROM markdowns
WHERE (relative_path, repo_id) IN (
  SELECT 
    unnest(@relative_path::text[]), 
    unnest(@repo_id::bigint[])
);
-- name: QueryMarkdown :many
select 
  users.username,r.repo_name, markdown_id,relative_path,users.user_id,r.repo_id,main_content,
  ROUND(ts_rank(fts_zh, plainto_tsquery($3))) + ROUND(ts_rank(fts_en, plainto_tsquery($3))) as rank,
  COALESCE(ts_headline(main_content,plainto_tsquery($3),'MaxFragments=10, MaxWords=7, MinWords=3'),'')
from markdowns 
JOIN repos as r on r.repo_id = markdowns.repo_id
JOIN users on users.user_id = r.user_id
where (fts_zh @@ plainto_tsquery($3) OR fts_en @@ plainto_tsquery($3))
  AND (
    (@role::text='admin' AND @signed::bool ) OR (
    users.blocked='false' AND (
      r.visibility_level = 'public'
      OR 
      (r.visibility_level = 'signed' AND @signed::bool)
      OR
      (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
      OR
      ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
    )
  )
  )
ORDER BY
  rank DESC
LIMIT $1
OFFSET $2;


-- name: QueryUserMarkdown :many
select 
  users.username,r.repo_name, markdown_id,relative_path,users.user_id,r.repo_id,main_content,
  ROUND(ts_rank(fts_zh, plainto_tsquery($4))) + ROUND(ts_rank(fts_en, plainto_tsquery($4))) as rank,
  COALESCE(ts_headline(main_content,plainto_tsquery($4),'MaxFragments=10, MaxWords=7, MinWords=3'),'')
from markdowns 
JOIN repos as r on r.repo_id = markdowns.repo_id
JOIN users on users.user_id = r.user_id
where users.user_id = $3  and (fts_zh @@ plainto_tsquery($4) OR fts_en @@ plainto_tsquery($4))
  AND (
    (@role::text='admin' AND @signed::bool ) OR (
    users.blocked='false' AND (
      r.visibility_level = 'public'
      OR 
      (r.visibility_level = 'signed' AND @signed::bool)
      OR
      (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
      OR
      ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
    )
  )
  )
ORDER BY
  rank DESC
LIMIT $1
OFFSET $2;

-- name: QueryRepoMarkdown :many
select 
  users.username,r.repo_name, markdown_id,relative_path,users.user_id,r.repo_id,main_content,
  ROUND(ts_rank(fts_zh, plainto_tsquery($4))) + ROUND(ts_rank(fts_en, plainto_tsquery($4))) as rank,
  COALESCE(ts_headline(main_content,plainto_tsquery($4),'MaxFragments=10, MaxWords=7, MinWords=3'),'')
from markdowns 
JOIN repos as r on r.repo_id = markdowns.repo_id
JOIN users on users.user_id = r.user_id
where users.user_id = $3 and r.repo_id = $5  and (fts_zh @@ plainto_tsquery($4) OR fts_en @@ plainto_tsquery($4))
ORDER BY
  rank DESC
LIMIT $1
OFFSET $2;


-- name: CreateMarkdown :one
INSERT INTO markdowns (
  relative_path,
  user_id,
  repo_id,
  main_content,
  table_content,
  md5,
  version_key
) VALUES (
  $1, $2, $3,$4,$5,$6,$7
) RETURNING *;

-- name: CreateMarkdownMulti :exec
INSERT INTO markdowns  (
  relative_path,
  user_id,
  repo_id,
  main_content,
  table_content,
  md5,
  version_key
)
SELECT unnest(@relative_path::text[]) AS relative_path,
  unnest(@user_id::bigint[]) AS user_id,
  unnest(@repo_id::bigint[]) AS repo_id,
  unnest(@main_content::text[]) AS main_content,
  unnest(@table_content::text[]) AS table_content,
  unnest(@md5::text[]) AS md5,
  unnest(@version_key::text[]) AS version_key;


-- name: QueryMd5ForCheck :many
SELECT relative_path,md5 FROM markdowns
WHERE repo_id = $1
ORDER BY relative_path ASC
FOR NO KEY UPDATE;

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



-- name: UpdateMarkdownVersionKey :exec
UPDATE markdowns AS m
SET version_key = tmp.version_key
FROM (
  SELECT 
  unnest(@version_key::text[]) AS version_key,
  unnest(@relative_path::text[]) AS relative_path,
  unnest(@repo_id::bigint[]) AS repo_id
) AS tmp
WHERE m.relative_path = tmp.relative_path and m.repo_id=tmp.repo_id;

-- name: UpdateMarkdownMulti :exec
UPDATE markdowns AS m
SET version_key = tmp.version_key,md5=tmp.md5,main_content=tmp.main_content,table_content=tmp.table_content,updated_at=now()
FROM (
  SELECT 
  unnest(@version_key::text[]) AS version_key,
  unnest(@relative_path::text[]) AS relative_path,
  unnest(@main_content::text[]) AS main_content,
  unnest(@table_content::text[]) AS table_content,
  unnest(@md5::text[]) AS md5,
  unnest(@repo_id::bigint[]) AS repo_id
) AS tmp
WHERE m.relative_path = tmp.relative_path and m.repo_id=tmp.repo_id;


-- name: DeleteMarkdown :exec
DELETE FROM markdowns
WHERE relative_path = $1 and repo_id = $2;


-- 每次更新仓库，都会给仓库一个新的版本key:version_key,有旧key的行都会被清理掉
-- name: DeleteOldMarkdown :exec
DELETE FROM markdowns
WHERE repo_id = $1 and version_key!=$2;

-- name: DeleteMarkdownByRepo :exec
DELETE FROM markdowns
WHERE repo_id = $1;


-- name: QueryUserAllMarkdown :many
select 
  users.username,repos.repo_name, markdown_id,relative_path,users.user_id,repos.repo_id,main_content,ts_rank(fts_zh, plainto_tsquery($4)) as rank,
  COALESCE(ts_headline(main_content,plainto_tsquery($4),'MaxFragments=10, MaxWords=7, MinWords=3'),'')
from markdowns 
JOIN repos on repos.repo_id = markdowns.repo_id
JOIN users on users.user_id = repos.user_id
where users.user_id = $3  and (fts_zh @@ plainto_tsquery($4) OR fts_en @@ plainto_tsquery($4))
ORDER BY
  rank DESC
LIMIT $1
OFFSET $2;


-- name: QueryUserVisibleMarkdown :many
select i.username,i.repo_name, i.markdown_id,i.relative_path,i.user_id,i.repo_id,COALESCE(ts_headline(i.main_content,plainto_tsquery($4),'MaxFragments=10, MaxWords=7, MinWords=3'),'')
from (
  select users.username,repos.repo_name, markdown_id,relative_path,users.user_id,repos.repo_id,main_content,ts_rank(fts_zh, plainto_tsquery($4)) as rank
  from markdowns 
  JOIN repos on repos.repo_id = markdowns.repo_id
  JOIN users on users.user_id = repos.user_id
  where markdowns.user_id = $3
        and (fts_zh @@ plainto_tsquery($4) OR fts_en @@ plainto_tsquery($4))
  ORDER BY
    rank DESC
  LIMIT $1
  OFFSET $2
) as i
where i.repo_id IN (SELECT repo_id FROM repos WHERE visibility_level = 'public')
ORDER BY
  i.rank DESC;

-- name: QueryRepoMarkdown :many
select i.username,i.repo_name, i.markdown_id,i.relative_path,i.user_id,i.repo_id,COALESCE(ts_headline(i.main_content,plainto_tsquery($5),'MaxFragments=10, MaxWords=7, MinWords=3'),'')
from (
  select repos.repo_name,users.username, markdown_id,relative_path,users.user_id,repos.repo_id,main_content,ts_rank(fts_zh, plainto_tsquery($5)) as rank
  from markdowns 
  JOIN repos on repos.repo_id = markdowns.repo_id
  JOIN users on users.user_id = repos.user_id
  where users.user_id = $3 and repos.repo_id=$4  and (fts_zh @@ plainto_tsquery($5) OR fts_en @@ plainto_tsquery($5))
  ORDER BY
    rank DESC
  LIMIT $1
  OFFSET $2
) as i
ORDER BY
  i.rank DESC;


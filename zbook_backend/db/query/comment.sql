-- name: CreateComment :one
INSERT INTO comments (
  user_id,
  repo_id,
  markdown_id,
  parent_id, 
  root_id,
  comment_content
) VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;


-- name: DeleteComment :exec
DELETE FROM comments
WHERE comment_id = $1;

-- name: GetCommentBasicInfo :one
SELECT comments.comment_id,comments.markdown_id,comments.user_id,comments.parent_id,comments.comment_content,comments.created_at,comments.root_id
FROM comments
WHERE comments.comment_id = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: GetCommentRepoInfo :one
SELECT repos.*
FROM markdowns
JOIN comments on markdowns.markdown_id=comments.markdown_id
JOIN repos on markdowns.repo_id = repos.repo_id
WHERE comments.comment_id = $1
LIMIT 1
FOR NO KEY UPDATE;


-- name: GetCommentDetail :one
SELECT comments.*,
  users.username,users.email,users.motto,users.created_at as user_created_at,
  COUNT(DISTINCT CASE WHEN comment_relations.relation_type = 'like' THEN comment_relations.relation_id END)  AS like_count,
  (SELECT COUNT(*) FROM comments c2 WHERE c2.root_id = comments.comment_id) AS reply_count,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id  and comment_relations.relation_type = 'like' and comment_relations.user_id = $2 ) as is_liked,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id  and comment_relations.relation_type = 'dislike' and comment_relations.user_id = $2 ) as is_disliked,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id  and comment_relations.relation_type = 'share' and comment_relations.user_id = $2 ) as is_shared,
  EXISTS(SELECT 1 FROM comment_reports WHERE comment_reports.comment_id = comments.comment_id and comment_reports.user_id = $2 ) as is_reported
FROM comments
LEFT JOIN comment_relations ON comments.comment_id = comment_relations.comment_id
JOIN users ON comments.user_id = users.user_id
WHERE comments.comment_id = $1
GROUP BY comments.comment_id,users.user_id
LIMIT 1;

-- name: ListCommentLevelOne :many
SELECT comments.*,
  users.username,users.email,users.motto,users.created_at as user_created_at,
  COUNT(DISTINCT CASE WHEN comment_relations.relation_type = 'like' THEN comment_relations.relation_id END)  AS like_count,
  (SELECT COUNT(*) FROM comments c2 WHERE c2.root_id = comments.comment_id) AS reply_count,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id and comment_relations.relation_type = 'like'  and comment_relations.user_id = $4 ) as is_liked,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id and comment_relations.relation_type = 'dislike'  and comment_relations.user_id = $4 ) as is_disliked,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id and comment_relations.relation_type = 'share'  and comment_relations.user_id = $4 ) as is_shared,
  EXISTS(SELECT 1 FROM comment_reports WHERE comment_reports.comment_id = comments.comment_id and comment_reports.user_id = $4 ) as is_reported
FROM comments
LEFT JOIN comment_relations ON comments.comment_id = comment_relations.comment_id
JOIN users ON comments.user_id = users.user_id
WHERE comments.markdown_id = $1 AND comments.parent_id IS NULL
GROUP BY comments.comment_id,users.user_id
ORDER BY comments.created_at DESC
LIMIT $2
OFFSET $3;


-- name: GetListCommentLevelOneCount :one
SELECT Count(*)
FROM comments
LEFT JOIN comment_relations ON comments.comment_id = comment_relations.comment_id
JOIN users ON comments.user_id = users.user_id
WHERE comments.markdown_id = $1 AND comments.parent_id IS NULL;


-- name: ListCommentLevelTwo :many
SELECT comments.*,
  users.username,users.email,users.motto,users.created_at as user_created_at, pu.username as pusername,
  COUNT(DISTINCT CASE WHEN comment_relations.relation_type = 'like' THEN comment_relations.relation_id END)  AS like_count,
  (SELECT COUNT(*) FROM comments c2 WHERE c2.root_id = comments.comment_id) AS reply_count,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id  and comment_relations.relation_type = 'like' and comment_relations.user_id = $4 ) as is_liked,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id  and comment_relations.relation_type = 'dislike' and comment_relations.user_id = $4 ) as is_disliked,
  EXISTS(SELECT 1 FROM comment_relations WHERE comment_relations.comment_id = comments.comment_id  and comment_relations.relation_type = 'share' and comment_relations.user_id = $4 ) as is_shared,
  EXISTS(SELECT 1 FROM comment_reports WHERE comment_reports.comment_id = comments.comment_id and comment_reports.user_id = $4 ) as is_reported
FROM comments
LEFT JOIN comments pc ON  comments.parent_id = pc.comment_id
LEFT JOIN users pu ON pu.user_id = pc.user_id
LEFT JOIN comment_relations ON comments.comment_id = comment_relations.comment_id
JOIN users ON comments.user_id = users.user_id
WHERE comments.root_id = $1
GROUP BY comments.comment_id,users.user_id,pu.username
ORDER BY comments.created_at
LIMIT $2
OFFSET $3;


-- name: GetListCommentLevelTwoCount :one
SELECT Count(*)
FROM comments
LEFT JOIN comments pc ON  comments.parent_id = pc.comment_id
LEFT JOIN users pu ON pu.user_id = pc.user_id
LEFT JOIN comment_relations ON comments.comment_id = comment_relations.comment_id
JOIN users ON comments.user_id = users.user_id
WHERE comments.root_id = $1;

-- name: ListComment :many
SELECT comments.*,
  users.username,users.email,users.created_at as user_created_at
FROM comments
JOIN markdowns on markdowns.markdown_id = comments.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users ON comments.user_id = users.user_id
JOIN users as mu ON mu.user_id=repos.user_id
ORDER BY comments.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetListCommentCount :one
SELECT COUNT(*)
FROM comments
JOIN markdowns on markdowns.markdown_id = comments.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users ON comments.user_id = users.user_id
JOIN users as mu ON mu.user_id=repos.user_id;

-- name: QueryComment :many
SELECT comments.*,
  ROUND(ts_rank(comments.fts_comment_zh, plainto_tsquery(@query))) + ROUND(ts_rank(comments.fts_comment_en, plainto_tsquery(@query))) as rank,
  users.username,users.email,users.created_at as user_created_at
FROM comments
JOIN markdowns on markdowns.markdown_id = comments.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users ON comments.user_id = users.user_id
JOIN users as mu ON mu.user_id=repos.user_id
WHERE (comments.fts_comment_zh @@ plainto_tsquery(@query) OR comments.fts_comment_en @@ plainto_tsquery(@query))
ORDER BY rank DESC
LIMIT $1
OFFSET $2;

-- name: GetQueryCommentCount :one
SELECT COUNT(*)
FROM comments
JOIN markdowns on markdowns.markdown_id = comments.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users ON comments.user_id = users.user_id
JOIN users as mu ON mu.user_id=repos.user_id
WHERE (comments.fts_comment_zh @@ plainto_tsquery(@query) OR comments.fts_comment_en @@ plainto_tsquery(@query));

-- name: CreateCommentRelation :exec
INSERT INTO comment_relations (
  user_id,
  comment_id,
  relation_type
) VALUES ($1,$2,$3);

-- name: DeleteCommentRelation :exec
DELETE FROM comment_relations
WHERE user_id=$1 and comment_id=$2 and relation_type=$3;

-- name: CreateCommentReport :exec
INSERT INTO comment_reports (
  user_id,
  comment_id,
  report_content
) VALUES ($1,$2,$3);


-- name: GetListCommentReportCount :one
SELECT 
  COUNT(*)
FROM comment_reports
JOIN users ON users.user_id = comment_reports.user_id
JOIN comments ON comments.comment_id = comment_reports.comment_id
JOIN markdowns ON comments.markdown_id = markdowns.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users ura ON ura.user_id = repos.user_id
JOIN users as uc ON comments.user_id = uc.user_id;

-- name: ListCommentReport :many
SELECT 
  comment_reports.*,users.username,comments.comment_content,
  repos.repo_name,ura.username as repo_username,markdowns.relative_path
FROM comment_reports
JOIN users ON users.user_id = comment_reports.user_id
JOIN comments ON comments.comment_id = comment_reports.comment_id
JOIN markdowns ON comments.markdown_id = markdowns.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users as ura ON ura.user_id = repos.user_id
JOIN users as uc ON comments.user_id = uc.user_id
ORDER BY comment_reports.created_at Desc
LIMIT $1
OFFSET $2;

-- name: QueryCommentReport :many
SELECT 
  comment_reports.*,ur.username,comments.comment_content,
    repos.repo_name,ura.username as repo_username,markdowns.relative_path,
      ROUND(ts_rank(comments.fts_comment_zh, plainto_tsquery(@query))) 
    + ROUND(ts_rank(comments.fts_comment_en, plainto_tsquery(@query))) 
    + ROUND(ts_rank(comment_reports.fts_report_zh, plainto_tsquery(@query)))
    + ROUND(ts_rank(comment_reports.fts_report_en, plainto_tsquery(@query)))
    + ROUND(ts_rank(ur.fts_username, plainto_tsquery(@query))) 
    + ROUND(ts_rank(uc.fts_username, plainto_tsquery(@query))) 
    + ROUND(ts_rank(repos.fts_repo_en, plainto_tsquery(@query))) 
    + ROUND(ts_rank(repos.fts_repo_zh, plainto_tsquery(@query))) 
     as rank
FROM comment_reports
JOIN users as ur ON ur.user_id = comment_reports.user_id
JOIN comments ON comments.comment_id = comment_reports.comment_id
JOIN markdowns ON comments.markdown_id = markdowns.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users as ura ON ura.user_id = repos.user_id
JOIN users as uc ON comments.user_id = uc.user_id
WHERE (
  comments.fts_comment_zh @@ plainto_tsquery(@query)
  OR comments.fts_comment_en @@ plainto_tsquery(@query)
  OR comment_reports.fts_report_zh @@ plainto_tsquery(@query) 
  OR comment_reports.fts_report_en @@ plainto_tsquery(@query) 
  OR uc.fts_username @@ plainto_tsquery(@query)  
  OR ur.fts_username @@ plainto_tsquery(@query)  
  OR repos.fts_repo_en @@ plainto_tsquery(@query)  
  OR repos.fts_repo_zh @@ plainto_tsquery(@query)  
  )
ORDER BY rank Desc
LIMIT $1
OFFSET $2;

-- name: GetQueryCommentReportCount :one
SELECT 
  COUNT(*)
FROM comment_reports
JOIN users as ur ON ur.user_id = comment_reports.user_id
JOIN comments ON comments.comment_id = comment_reports.comment_id
JOIN markdowns ON comments.markdown_id = markdowns.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users as uc ON comments.user_id = uc.user_id
WHERE (
  comments.fts_comment_zh @@ plainto_tsquery(@query)
  OR comments.fts_comment_en @@ plainto_tsquery(@query)
  OR comment_reports.fts_report_zh @@ plainto_tsquery(@query) 
  OR comment_reports.fts_report_en @@ plainto_tsquery(@query) 
  OR uc.fts_username @@ plainto_tsquery(@query)  
  OR ur.fts_username @@ plainto_tsquery(@query)  
  );
-- name: UpdateCommentReportStatus :exec
UPDATE comment_reports
SET processed=$2
WHERE report_id = $1;

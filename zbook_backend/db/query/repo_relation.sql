-- name: CreateRepoRelation :exec
INSERT INTO repo_relations (
  user_id,
  repo_id,
  relation_type
) VALUES ($1,$2,$3);

-- name: DeleteRepoRelation :exec
DELETE FROM repo_relations
WHERE user_id=$1 and repo_id=$2 and relation_type = $3;

-- name: GetRepoRelation :one
SELECT *
FROM repo_relations
WHERE user_id = $1 and repo_id=$2 and relation_type = $3;


-- name: GetListSelectedUserByRepoCount :one
SELECT COUNT(*)
FROM repos as r
LEFT JOIN repo_relations as rr ON rr.repo_id=r.repo_id
JOIN users as u ON u.user_id = rr.user_id AND (u.blocked='false' OR @role::text='admin')
WHERE r.repo_id=$1 AND rr.relation_type = 'visi';

-- name: ListSelectedUserByRepo :many
SELECT u.*
FROM repos as r
LEFT JOIN repo_relations as rr ON rr.repo_id=r.repo_id
JOIN users as u ON u.user_id = rr.user_id AND (u.blocked='false' OR @role::text='admin')
WHERE r.repo_id=$3 AND rr.relation_type = 'visi'
ORDER BY rr.created_at DESC
LIMIT $1
OFFSET $2;


-- name: QuerySelectedUserByRepo :many
select u.*,ts_rank(fts_username, plainto_tsquery(@query)) as rank
FROM repos as r
LEFT JOIN repo_relations as rr ON rr.repo_id=r.repo_id
JOIN users as u ON u.user_id = rr.user_id
WHERE r.repo_id=$3 AND rr.relation_type = 'visi' AND (u.blocked='false' OR @role::text='admin') AND fts_username @@ plainto_tsquery(@query)
ORDER BY rank DESC
LIMIT $1
OFFSET $2;

-- name: GetQuerySelectedUserByRepoCount :one
SELECT COUNT(*)
FROM repos as r
LEFT JOIN repo_relations as rr ON rr.repo_id=r.repo_id
JOIN users as u ON u.user_id = rr.user_id
WHERE r.repo_id=$1 AND rr.relation_type = 'visi' AND (u.blocked='false' OR @role::text='admin') AND fts_username @@ plainto_tsquery(@query);


-- name: QueryUserByRepo :many
SELECT
   u.*,ts_rank(fts_username, plainto_tsquery(@query)) as rank,
   CASE 
       WHEN MAX(CASE WHEN rr.relation_type = 'visi' THEN rr.user_id END) IS NOT NULL 
       THEN true 
       ELSE false 
   END AS is_visible
FROM 
  users as u 
LEFT JOIN 
    repo_relations rr ON rr.user_id = u.user_id AND rr.repo_id=$3
WHERE (u.blocked='false' OR @role::text='admin') AND fts_username @@ plainto_tsquery(@query)
GROUP BY u.user_id,rr.created_at
ORDER BY rank DESC,rr.created_at DESC
LIMIT $1
OFFSET $2;
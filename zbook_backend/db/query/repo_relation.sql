-- name: CreateRepoRelation :exec
INSERT INTO repo_relations (
  user_id,
  repo_id,
  relation_type
) VALUES ($1,$2,$3);

-- name: DeleteRepoRelation :exec
DELETE FROM repo_relations
WHERE user_id=$1 and repo_id=$2 and relation_type = $3;

-- name: CreateRepoVisibility :exec
INSERT INTO repo_visibility (
  repo_id,
  user_id
) VALUES ($1,$2);

-- name: DeleteRepoVisibility :exec
DELETE FROM repo_visibility
WHERE user_id=$1 and repo_id=$2;


-- name: GetRepoVisibility :one
SELECT *
FROM repos
WHERE user_id = $1 and repo_id=$2;

-- name: ListRepoVisibilityByRepo :many
SELECT u.*
FROM repos as r
LEFT JOIN repo_visibility as rv ON rv.repo_id=r.repo_id
JOIN users as u ON u.user_id = rv.user_id
WHERE r.repo_id=$3
ORDER BY u.user_id
LIMIT $1
OFFSET $2;

-- name: GetRepoVisibilityByRepoCount :one
SELECT COUNT(*)
FROM repos as r
LEFT JOIN repo_visibility as rv ON rv.repo_id=r.repo_id
JOIN users as u ON u.user_id = rv.user_id
WHERE r.repo_id=$1;


-- name: QueryRepoVisibilityByRepo :many
SELECT
   u.*,
   CASE WHEN MAX(rv.user_id) IS NOT NULL THEN true ELSE false END AS is_visible
FROM 
  users as u 
LEFT JOIN 
    repo_visibility rv ON rv.user_id = u.user_id AND rv.repo_id=$3
WHERE u.username=$4
GROUP BY u.user_id
ORDER BY u.user_id
LIMIT $1
OFFSET $2;
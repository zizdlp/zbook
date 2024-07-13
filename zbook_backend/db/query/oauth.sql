-- name: CreateOAuth :one
INSERT INTO oauths (
  user_id,
  oauth_type,
  app_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetOAuthUser :one
SELECT *
FROM oauths
JOIN users ON oauths.user_id = users.user_id
WHERE oauth_type = $1 and app_id = $2 LIMIT 1
FOR NO KEY UPDATE;

-- name: CheckOAuthStatus :one
SELECT
  COALESCE(COUNT(CASE WHEN "oauth_type" = 'github' THEN 1 END), 0) > 0 AS github_status,
  COALESCE(COUNT(CASE WHEN "oauth_type" = 'google' THEN 1 END), 0) > 0 AS google_status
FROM
  "oauths"
WHERE
  "user_id" = $1;

-- name: DeleteOAuth :one
DELETE FROM oauths
WHERE user_id=$1 and oauth_type=$2
RETURNING *;

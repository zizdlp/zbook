-- name: GetUserByID :one
SELECT *
FROM users
WHERE user_id = $1 
LIMIT 1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1 
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE users.email = $1 
LIMIT 1;

-- name: GetUnReadCount :one
SELECT unread_count
FROM users
WHERE username = $1 
LIMIT 1;

-- name: ListUser :many
SELECT *
FROM users u
WHERE @signed::bool AND u.deleted='false' AND (@role::text='admin' OR u.blocked='false')
ORDER BY u.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetListUserCount :one
SELECT COUNT(*)
FROM users u
WHERE @signed::bool AND u.deleted='false' AND (@role::text='admin' OR u.blocked='false');

-- name: QueryUser :many
select users.*,ts_rank(fts_username, plainto_tsquery(@query)) as rank
from users 
where fts_username @@ plainto_tsquery(@query) AND @signed::bool AND users.deleted='false' AND (@role::text='admin' OR users.blocked='false')
ORDER BY rank DESC
LIMIT $1
OFFSET $2;

-- name: GetQueryUserCount :one
select COUNT(*)
from users 
where fts_username @@ plainto_tsquery(@query) AND @signed::bool AND users.deleted='false' AND (@role::text='admin' OR users.blocked='false');

-- name: GetDailyCreateUserCount :many
SELECT DATE(created_at) AS registration_date, COUNT(*) AS new_users_count
FROM users
WHERE created_at >= CURRENT_DATE - INTERVAL '7 days'
GROUP BY registration_date
ORDER BY registration_date;


-- name: GetUserInfo :one
WITH liked_repos_count AS (
  SELECT
    Count(*) as like_count
  FROM
      repos r
  JOIN 
      repo_relations AS rr ON r.repo_id = rr.repo_id
  JOIN
      users as ur ON ur.user_id=r.user_id
  JOIN
      users as uq ON uq.user_id=rr.user_id
  WHERE
    uq.user_id = @user_id AND rr.relation_type='like' AND uq.deleted = FALSE AND ur.deleted=FALSE  AND ( 
      (@role::text='admin' AND @signed::bool ) OR (
        uq.blocked = FALSE AND ur.blocked =FALSE AND 
        (
          (r.visibility_level = 'public' ) 
          OR
          (r.visibility_level = 'signed' AND @signed::bool) 
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_visibility WHERE repo_visibility.repo_id = r.repo_id AND repo_visibility.user_id = @cur_user_id))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    )
),
 owned_repos_count AS (
  SELECT
    COUNT(*) as repo_count
  FROM
      repos r
  JOIN
      users as u ON u.user_id=r.user_id
  WHERE
      u.user_id = @user_id AND u.deleted = FALSE AND (
        (@role::text='admin' AND @signed::bool ) OR (
          u.blocked='false' AND (
            r.visibility_level = 'public'
            OR 
            (r.visibility_level = 'signed' AND @signed::bool)
            OR
            (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_visibility WHERE repo_visibility.repo_id = r.repo_id AND repo_visibility.user_id = @cur_user_id))
            OR
            ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
          )
        )
      )
)
SELECT 
    u.*,
    repo_count,
    like_count,
    (SELECT COUNT(*) FROM follows f1 JOIN users as uf ON uf.user_id = f1.follower_id WHERE f1.following_id = u.user_id and uf.deleted = 'false' and (uf.blocked = 'false' OR @role::text='admin')) AS follower_count,
    (SELECT COUNT(*) FROM follows f2 JOIN users as uf ON uf.user_id = f2.following_id WHERE f2.follower_id = u.user_id and uf.deleted = 'false' and (uf.blocked = 'false' OR @role::text='admin')) AS following_count,
    EXISTS(SELECT 1 FROM follows WHERE follows.follower_id = @cur_user_id AND follows.following_id = @user_id) AS is_following
FROM users u
JOIN liked_repos_count lrc ON 1=1
JOIN owned_repos_count ownrc ON 1=1
WHERE u.user_id = @user_id;

-- name: CreateSession :one
INSERT INTO sessions (
  session_id,
  user_id,
  refresh_token,
  user_agent,
  client_ip,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE session_id = $1 LIMIT 1;

-- name: GetListActiveSessionCount :one
SELECT Count(*) FROM sessions
WHERE sessions.expires_at > NOW() LIMIT 1;

-- name: ListActiveSession :many
SELECT 
  * 
FROM 
  sessions
INNER JOIN users ON users.user_id = sessions.user_id
WHERE sessions.expires_at > NOW()
ORDER BY sessions.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetQueryActiveSessionCount :one
select Count(*)
FROM 
  sessions
JOIN users ON users.user_id = sessions.user_id
WHERE sessions.expires_at > NOW() AND fts_username @@ plainto_tsquery(@query);

-- name: QueryActiveSession :many
select sessions.*,ts_rank(users.fts_username, plainto_tsquery(@query)) as rank,users.*
FROM 
  sessions
JOIN users ON users.user_id = sessions.user_id
WHERE sessions.expires_at > NOW() AND fts_username @@ plainto_tsquery(@query)
ORDER BY
  rank DESC
LIMIT $1
OFFSET $2;

-- name: GetDailyActiveUserCount :many
SELECT DATE(created_at) AS registration_date, COUNT(DISTINCT user_id) AS active_users_count
FROM sessions
WHERE created_at >= CURRENT_DATE - INTERVAL '7 days'
GROUP BY registration_date
ORDER BY registration_date;

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

-- name: GetListSessionCount :one
SELECT Count(*) FROM sessions;

-- name: ListSession :many
SELECT 
  * 
FROM 
  sessions
INNER JOIN users ON users.user_id = sessions.user_id
ORDER BY sessions.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetQuerySessionCount :one
select Count(*)
FROM 
  sessions
JOIN users ON users.user_id = sessions.user_id
WHERE fts_username @@ plainto_tsquery(@query);

-- name: QuerySession :many
SELECT 
  sessions.*,
  ts_rank(users.fts_username, plainto_tsquery(@query)) as rank,
  users.*
FROM 
  sessions
JOIN users ON users.user_id = sessions.user_id
WHERE fts_username @@ plainto_tsquery(@query)
ORDER BY
  rank DESC,
  sessions.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetDailyActiveUserCount :many
SELECT (created_at AT TIME ZONE @timezone)::date AS registration_date, COUNT(DISTINCT user_id) AS active_users_count
FROM sessions
WHERE  (created_at AT TIME ZONE @timezone) >= (CURRENT_DATE AT TIME ZONE @timezone) - (@interval_days || ' days')::INTERVAL
GROUP BY registration_date
ORDER BY registration_date DESC;
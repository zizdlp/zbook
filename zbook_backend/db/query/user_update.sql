-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  hashed_password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

-- name: UpdateUnreadCount :exec
UPDATE users
SET unread_count = (
  SELECT COUNT(*) FROM (
    SELECT noti_id FROM follower_notifications
    WHERE user_id = users.user_id AND created_at  > users.unread_count_updated_at AND readed = false 
    UNION ALL
    SELECT noti_id FROM comment_notifications
    WHERE user_id = users.user_id AND created_at  > users.unread_count_updated_at AND readed = false 
    UNION ALL
    SELECT noti_id FROM system_notifications
    WHERE user_id = users.user_id AND created_at  > users.unread_count_updated_at AND readed = false 
    UNION ALL
    SELECT noti_id FROM repo_notifications
    WHERE user_id = users.user_id AND created_at  > users.unread_count_updated_at AND readed = false 
  ) AS subquery
) 
WHERE users.user_id = $1;

-- name: ResetUnreadCount :exec
UPDATE users
SET unread_count_updated_at = now(),
  unread_count = (
  SELECT COUNT(*) FROM (
    SELECT noti_id FROM follower_notifications
    WHERE user_id = users.user_id AND created_at  > now() AND readed = false 
    UNION ALL
    SELECT noti_id FROM comment_notifications
    WHERE user_id = users.user_id AND created_at  > now() AND readed = false 
    UNION ALL
    SELECT noti_id FROM system_notifications
    WHERE user_id = users.user_id AND created_at  > now() AND readed = false 
    UNION ALL
    SELECT noti_id FROM repo_notifications
    WHERE user_id = users.user_id AND created_at  > now() AND readed = false 
  ) AS subquery
) 
WHERE users.username = $1;

-- name: UpdateUserBasicInfo :one
UPDATE users
SET motto=COALESCE(sqlc.narg(motto),motto),
hashed_password=COALESCE(sqlc.narg(hashed_password),hashed_password),
user_role=COALESCE(sqlc.narg(user_role),user_role),
onboarding=COALESCE(sqlc.narg(onboarding),onboarding),
blocked=COALESCE(sqlc.narg(blocked),blocked),
verified=COALESCE(sqlc.narg(verified),verified)
WHERE username = sqlc.arg(username)
RETURNING *;

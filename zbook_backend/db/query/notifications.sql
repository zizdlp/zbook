-- name: CreateSystemNotification :one
INSERT INTO system_notifications (
    user_id,
    title,
    contents,
    redirect_url
) VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: ListSystemNotification :many
SELECT system_notifications.noti_id,system_notifications.readed,system_notifications.created_at,
       title,contents,redirect_url
FROM system_notifications
WHERE user_id=$1 AND (system_notifications.readed = 'false' OR system_notifications.created_at >= NOW() - INTERVAL '14 days')
ORDER BY system_notifications.created_at Desc
LIMIT $2
OFFSET $3;

-- name: GetListSystemNotificationUnReadedCount :one
SELECT Count(*)
FROM system_notifications
WHERE user_id=$1 AND (system_notifications.readed = 'false');

-- name: MarkSystemNotificationReaded :one
UPDATE system_notifications
SET readed = 'true'
WHERE noti_id = $1 AND user_id = $2
RETURNING *;

-- name: CreateFollowerNotification :one
INSERT INTO follower_notifications (
    user_id,
    follower_id
) VALUES ($1,$2)
RETURNING *;

-- name: DeleteFollowerNotification :one
DELETE FROM follower_notifications
WHERE user_id=$1 and follower_id=$2
RETURNING *;

-- name: MarkFollowerNotificationReaded :one
UPDATE follower_notifications
SET readed = 'true'
WHERE noti_id=$1 and user_id = $2
RETURNING *;

-- name: ListFollowerNotification :many
SELECT 
    users.username,users.email,
    follower_notifications.readed,follower_notifications.noti_id,follower_notifications.created_at
FROM 
    users
JOIN
    follower_notifications ON users.user_id = follower_notifications.follower_id
WHERE 
    follower_notifications.user_id=$1 AND (follower_notifications.readed = 'false' OR follower_notifications.created_at >= NOW() - INTERVAL '14 days')
ORDER BY 
    follower_notifications.created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetListFollowerNotificationUnreadedCount :one
SELECT 
   Count(*)
FROM 
    users
JOIN
    follower_notifications ON users.user_id = follower_notifications.follower_id
WHERE 
    follower_notifications.user_id=$1 AND (follower_notifications.readed = 'false');

-- name: CreateCommentNotification :one
INSERT INTO comment_notifications (
    user_id,
    comment_id
) VALUES ($1,$2)
RETURNING *;

-- name: MarkCommentNotificationReaded :one
UPDATE comment_notifications
SET readed = 'true'
WHERE noti_id=$1 and user_id = $2
RETURNING *;


-- name: ListCommentNotification :many
SELECT 
    users.username,users.email,ru.username as repo_username,
    comment_notifications.readed, comment_notifications.noti_id,comment_notifications.created_at,
    markdowns.repo_id,markdowns.relative_path,comments.comment_content,repos.repo_name
FROM 
    users
JOIN 
    comments ON comments.user_id=users.user_id
JOIN 
    markdowns ON markdowns.markdown_id=comments.markdown_id
JOIN repos ON repos.repo_id = markdowns.repo_id
JOIN users as ru ON repos.user_id = ru.user_id
JOIN 
    comment_notifications ON comment_notifications.comment_id=comments.comment_id
WHERE 
    comment_notifications.user_id=$1 AND (comment_notifications.readed = 'false' OR comment_notifications.created_at >= NOW() - INTERVAL '14 days')
ORDER BY 
    comment_notifications.created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetListCommentNotificationUnreadedCount :one
SELECT 
  Count(*)
FROM 
    users
JOIN 
    comments ON comments.user_id=users.user_id
JOIN 
    markdowns ON markdowns.markdown_id=comments.markdown_id
JOIN 
    comment_notifications ON comment_notifications.comment_id=comments.comment_id
WHERE 
    comment_notifications.user_id=$1 AND (comment_notifications.readed = 'false');

-- name: CreateRepoNotification :one
INSERT INTO repo_notifications (
    user_id,
    repo_id
) VALUES ($1,$2)
RETURNING *;

-- name: MarkRepoNotificationReaded :one
UPDATE repo_notifications
SET readed = 'true'
WHERE noti_id=$1 and user_id = $2
RETURNING *;


-- name: ListRepoNotification :many
SELECT 
    users.username,users.email,
    repo_notifications.readed,repo_notifications.noti_id,repo_notifications.created_at,
    repos.repo_id,repos.repo_name
FROM repo_notifications
JOIN repos ON repos.repo_id = repo_notifications.repo_id
JOIN users ON users.user_id = repos.user_id
WHERE repo_notifications.user_id=$1  AND (repo_notifications.readed = 'false' OR repo_notifications.created_at >= NOW() - INTERVAL '14 days')
ORDER BY repo_notifications.created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetListRepoNotificationUnreadedCount :one
SELECT 
   Count(*)
FROM repo_notifications
JOIN repos ON repos.repo_id = repo_notifications.repo_id
JOIN users ON users.user_id = repos.user_id
WHERE repo_notifications.user_id=$1  AND (repo_notifications.readed = 'false' );



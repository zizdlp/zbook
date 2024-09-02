-- name: CreateRepo :one
INSERT INTO repos (
  user_id,
  git_protocol,
  git_host,
  git_username,
  git_repo,
  git_access_token,
  repo_name,
  theme_sidebar,
  theme_color,
  repo_description,
  sync_token,
  commit_id,
  visibility_level
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) 
RETURNING *;

-- name: UpdateRepoConfig :exec
UPDATE repos
SET config=$2,commit_id=$3,home=$4,updated_at=now()
WHERE repo_id = $1;

-- name: UpdateRepoInfo :one
UPDATE repos
SET 
repo_name=COALESCE(sqlc.narg(repo_name),repo_name),
repo_description=COALESCE(sqlc.narg(repo_description),repo_description),
sync_token=COALESCE(sqlc.narg(sync_token),sync_token),
visibility_level=COALESCE(sqlc.narg(visibility_level),visibility_level),
git_access_token=COALESCE(sqlc.narg(git_access_token),git_access_token),
theme_sidebar=COALESCE(sqlc.narg(theme_sidebar),theme_sidebar),
theme_color=COALESCE(sqlc.narg(theme_color),theme_color)
WHERE repo_id = sqlc.arg(repo_id)
RETURNING *;

-- name: DeleteRepo :exec
DELETE FROM repos
WHERE repo_id = $1;

-- name: GetRepo :one
SELECT * from repos
WHERE repo_id = $1;

-- name: GetRepoID :one
SELECT repos.repo_id 
from repos
JOIN users on users.user_id= repos.user_id
WHERE users.username=$1 AND repos.repo_name=$2;

-- name: GetRepoByRepoName :one
SELECT * from repos
JOIN users on users.user_id= repos.user_id
WHERE users.username=$1 AND repos.repo_name=$2;

-- name: GetRepoConfig :one
SELECT repos.repo_id,config,repos.user_id,visibility_level,repos.theme_sidebar,repos.theme_color,repos.home FROM repos
JOIN users on users.user_id = repos.user_id
WHERE users.username=$1 AND repos.repo_name=$2;

-- name: GetRepoHome :one
SELECT repos.home FROM repos
JOIN users on users.user_id = repos.user_id
WHERE users.username=$1 AND repos.repo_name=$2;


-- name: GetRepoPermission :one
SELECT 
  repos.visibility_level as visibility_level,
  users.user_id,users.blocked as user_blocked,users.username,
  users.user_role as user_role,
  repos.repo_id
FROM 
  repos
INNER JOIN users ON users.user_id = repos.user_id
WHERE 
  repo_id = $1;

-- name: GetRepoBasicInfo :one
SELECT repos.*,
  users.username, users.email
FROM repos
INNER JOIN users ON repos.user_id = users.user_id
WHERE users.username=$1 AND repos.repo_name=$2;


-- name: GetQueryRepoCount :one
SELECT
   count(*)
FROM
    repos r
JOIN 
  users u ON u.user_id = r.user_id
where (r.fts_repo_en @@ plainto_tsquery(@query) OR r.fts_repo_zh @@ plainto_tsquery(@query))
  AND (
    (@role::text='admin' AND @signed::bool ) OR (
      u.blocked='false' AND (
        r.visibility_level = 'public'
        OR 
        (r.visibility_level = 'signed' AND @signed::bool)
        OR
        (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
        OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
      )
    )
  );
  
-- name: QueryRepo :many
select 
  r.*,
  u.username,
  ROUND(ts_rank(r.fts_repo_en, plainto_tsquery(@query))) + ROUND(ts_rank(r.fts_repo_zh, plainto_tsquery(@query))) as rank
FROM
    repos r
JOIN 
  users u ON u.user_id = r.user_id
where (r.fts_repo_en @@ plainto_tsquery(@query) OR r.fts_repo_zh @@ plainto_tsquery(@query))
  AND (
    (@role::text='admin' AND @signed::bool ) OR (
      u.blocked='false' AND (
        r.visibility_level = 'public'
        OR 
        (r.visibility_level = 'signed' AND @signed::bool)
        OR
        (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
        OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
      )
    )
  )
ORDER BY rank DESC
LIMIT $1
OFFSET $2;


-- name: GetListRepoCount :one
SELECT
   count(*)
FROM
    repos r
JOIN 
  users u ON u.user_id = r.user_id
WHERE
  (@role::text='admin' AND @signed::bool ) OR (
      u.blocked='false' AND (
        r.visibility_level = 'public'
        OR 
        (r.visibility_level = 'signed' AND @signed::bool)
        OR
        (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
        OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
      )
  );

-- name: ListRepo :many
SELECT
   r.*,
   (SELECT COUNT(*) FROM repo_relations WHERE repo_id = r.repo_id and relation_type = 'like') AS like_count,
   u.username,
   EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id  and repo_relations.relation_type = 'like' and repo_relations.user_id = @cur_user_id ) as is_liked
FROM
    repos r
JOIN 
  users u ON u.user_id = r.user_id
WHERE
  (@role::text='admin' AND @signed::bool ) OR (
    u.blocked='false' AND (
      r.visibility_level = 'public'
      OR 
      (r.visibility_level = 'signed' AND @signed::bool)
      OR
      (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
      OR
      ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
    )
  )
    
ORDER BY r.created_at DESC
LIMIT $1
OFFSET $2;


-- name: GetQueryUserOwnRepoCount :one
SELECT
  COUNT(*)
FROM
    repos r
JOIN
    users as u ON u.user_id=r.user_id
WHERE
    (r.fts_repo_en @@ plainto_tsquery(@query) OR r.fts_repo_zh @@ plainto_tsquery(@query)) AND u.user_id = @user_id AND (
      (@role::text='admin' AND @signed::bool ) OR (
        u.blocked='false' AND (
          r.visibility_level = 'public'
          OR 
          (r.visibility_level = 'signed' AND @signed::bool)
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    );

-- name: QueryUserOwnRepo :many
SELECT
  r.*,
  ROUND(ts_rank(r.fts_repo_en, plainto_tsquery(@query))) + ROUND(ts_rank(r.fts_repo_zh, plainto_tsquery(@query))) as rank,
  (SELECT COUNT(*) FROM repo_relations WHERE repo_id = r.repo_id and relation_type = 'like') AS like_count,
  EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id  and repo_relations.relation_type = 'like' and repo_relations.user_id = @cur_user_id ) as is_liked
FROM
    repos r
JOIN
    users as u ON u.user_id=r.user_id
WHERE
    (r.fts_repo_en @@ plainto_tsquery(@query) OR r.fts_repo_zh @@ plainto_tsquery(@query)) AND u.user_id = @user_id AND (
      (@role::text='admin' AND @signed::bool ) OR (
        u.blocked='false' AND (
          r.visibility_level = 'public'
          OR 
          (r.visibility_level = 'signed' AND @signed::bool)
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    )
ORDER BY rank DESC
LIMIT $1
OFFSET $2;

-- name: ListUserOwnRepo :many
SELECT
   r.*,
  (SELECT COUNT(*) FROM repo_relations WHERE repo_id = r.repo_id and relation_type = 'like') AS like_count,
  EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.relation_type = 'like' AND repo_relations.user_id = @cur_user_id) AS is_liked
FROM
    repos r
JOIN
    users as u ON u.user_id=r.user_id
WHERE
    u.user_id = @user_id AND (
      (@role::text='admin' AND @signed::bool ) OR (
        u.blocked='false' AND (
          r.visibility_level = 'public'
          OR 
          (r.visibility_level = 'signed' AND @signed::bool)
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    )
ORDER BY r.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetListUserOwnRepoCount :one
SELECT
    COUNT(*)
FROM
    repos r
JOIN
    users as u ON u.user_id=r.user_id
WHERE
    u.user_id = @user_id AND (
      (@role::text='admin' AND @signed::bool ) OR (
        u.blocked='false' AND (
          r.visibility_level = 'public'
          OR 
          (r.visibility_level = 'signed' AND @signed::bool)
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    );

-- name: QueryUserLikeRepo :many
SELECT
  r.*,ur.username,
  ROUND(ts_rank(r.fts_repo_en, plainto_tsquery(@query))) + ROUND(ts_rank(r.fts_repo_zh, plainto_tsquery(@query))) as rank,
  (SELECT COUNT(*) FROM repo_relations WHERE repo_id = r.repo_id and relation_type = 'like') AS like_count,
    EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id  and repo_relations.relation_type = 'like' and repo_relations.user_id = @cur_user_id ) as is_liked
FROM
    repos r
JOIN repo_relations AS rr ON r.repo_id = rr.repo_id
JOIN
    users as ur ON ur.user_id=r.user_id
JOIN
    users as uq ON uq.user_id=rr.user_id
WHERE
    (r.fts_repo_en @@ plainto_tsquery(@query) OR r.fts_repo_zh @@ plainto_tsquery(@query)) AND uq.user_id = @user_id AND rr.relation_type='like' AND ( 
      (@role::text='admin' AND @signed::bool ) OR (
        uq.blocked = FALSE AND ur.blocked =FALSE AND 
        (
          (r.visibility_level = 'public' ) 
          OR
          (r.visibility_level = 'signed' AND @signed::bool) 
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    )
ORDER BY rank DESC
LIMIT $1
OFFSET $2;

-- name: GetQueryUserLikeRepoCount :one
SELECT
  COUNT(*)
FROM
    repos r
JOIN repo_relations AS rr ON r.repo_id = rr.repo_id
JOIN
    users as ur ON ur.user_id=r.user_id
JOIN
    users as uq ON uq.user_id=rr.user_id
WHERE
    (r.fts_repo_en @@ plainto_tsquery(@query) OR r.fts_repo_zh @@ plainto_tsquery(@query)) AND uq.user_id = @user_id AND rr.relation_type='like'  AND ( 
      (@role::text='admin' AND @signed::bool ) OR (
        uq.blocked = FALSE AND ur.blocked =FALSE AND 
        (
          (r.visibility_level = 'public' ) 
          OR
          (r.visibility_level = 'signed' AND @signed::bool) 
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    );

-- name: ListUserLikeRepo :many
SELECT
   r.*,ur.username,
  (SELECT COUNT(*) FROM repo_relations WHERE repo_id = r.repo_id and relation_type = 'like') AS like_count,
  EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id  and repo_relations.relation_type = 'like' and repo_relations.user_id = @cur_user_id ) as is_liked
FROM
    repos r
JOIN 
    repo_relations AS rr ON r.repo_id = rr.repo_id
JOIN
    users as ur ON ur.user_id=r.user_id -- query user likes repo owner
JOIN
    users as uq ON uq.user_id=rr.user_id -- query user
WHERE
    uq.user_id = @user_id AND rr.relation_type='like' AND ( 
      (@role::text='admin' AND @signed::bool ) OR (
        uq.blocked = FALSE AND ur.blocked =FALSE AND 
        (
          (r.visibility_level = 'public' ) 
          OR
          (r.visibility_level = 'signed' AND @signed::bool) 
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    )
ORDER BY r.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetListUserLikeRepoCount :one
SELECT
 COUNT(*)
FROM
    repos r
JOIN repo_relations AS rr ON r.repo_id = rr.repo_id
JOIN
    users as ur ON ur.user_id=r.user_id
JOIN
    users as uq ON uq.user_id=rr.user_id
WHERE
    uq.user_id = @user_id AND rr.relation_type='like' AND ( 
      (@role::text='admin' AND @signed::bool ) OR (
        uq.blocked = FALSE AND ur.blocked =FALSE AND 
        (
          (r.visibility_level = 'public' ) 
          OR
          (r.visibility_level = 'signed' AND @signed::bool) 
          OR
          (r.visibility_level = 'chosen' AND @signed::bool AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi'))
          OR
          ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND r.user_id = @cur_user_id AND @signed::bool)
        )
      )
    );
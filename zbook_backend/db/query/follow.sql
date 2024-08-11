-- name: CreateFollow :one
INSERT INTO follows (
  follower_id,
  following_id
) VALUES ($1, $2) 
RETURNING *;

-- name: IsFollowing :one
SELECT EXISTS (
  SELECT 1
  FROM follows
  WHERE follower_id = $1
    AND following_id = $2
  LIMIT 1
);

-- name: DeleteFollow :one
DELETE FROM follows
WHERE follower_id= $1 and following_id=$2
RETURNING follow_id;

-- name: ListFollower :many
SELECT 
    u.*,
    CASE WHEN MAX(ff.follower_id) IS NOT NULL THEN true ELSE false END AS is_following,
    COUNT(DISTINCT r.repo_id) as repo_count
FROM 
    users u
JOIN 
    follows f ON f.follower_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.following_id = @user_id AND (u.blocked='false' OR @role::text='admin')
GROUP BY 
    u.user_id,f.created_at
ORDER BY 
    f.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetListFollowerCount :one
SELECT 
    COUNT(*)
FROM 
    users u
JOIN 
    follows f ON f.follower_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.following_id = @user_id AND (u.blocked='false' OR @role::text='admin');

-- name: QueryFollower :many
SELECT 
    u.*,
    ts_rank(u.fts_username, plainto_tsquery(@query)) as rank,
    CASE WHEN MAX(ff.follower_id) IS NOT NULL THEN true ELSE false END AS is_following,
    COUNT(DISTINCT r.repo_id) as repo_count
FROM 
    users u
JOIN 
    follows f ON f.follower_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.following_id = @user_id and u.fts_username @@ plainto_tsquery(@query) AND (u.blocked='false' OR @role::text='admin')
GROUP BY 
    u.user_id
ORDER BY 
    rank DESC
LIMIT $1
OFFSET $2;


-- name: GetQueryFollowerCount :one
SELECT 
    COUNT(*)
FROM 
    users u
JOIN 
    follows f ON f.follower_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
       ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.following_id = @user_id and u.fts_username @@ plainto_tsquery(@query) AND (u.blocked='false' OR @role::text='admin');

-- name: ListFollowing :many
SELECT 
    u.*,
    CASE WHEN MAX(ff.follower_id) IS NOT NULL THEN true ELSE false END AS is_following,
    COUNT(DISTINCT r.repo_id) as repo_count
FROM 
    users u
JOIN 
    follows f ON f.following_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
       ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.follower_id = @user_id AND (u.blocked='false' OR @role::text='admin')
GROUP BY 
    u.user_id,f.created_at
ORDER BY 
    f.created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetListFollowingCount :one
SELECT 
    COUNT(*)
FROM 
    users u
JOIN 
    follows f ON f.following_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.follower_id = @user_id AND (u.blocked='false' OR @role::text='admin');



-- name: QueryFollowing :many
SELECT 
    u.*,
    ts_rank(u.fts_username, plainto_tsquery(@query)) as rank,
    CASE WHEN MAX(ff.follower_id) IS NOT NULL THEN true ELSE false END AS is_following,
    COUNT(DISTINCT r.repo_id) as repo_count
FROM 
    users u
JOIN 
    follows f ON f.following_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.follower_id = @user_id and u.fts_username @@ plainto_tsquery(@query) AND (u.blocked='false' OR @role::text='admin')
GROUP BY 
    u.user_id
ORDER BY 
    rank DESC
LIMIT $1
OFFSET $2;



-- name: GetQueryFollowingCount :one
SELECT 
    COUNT(*)
FROM 
    users u
JOIN 
    follows f ON f.following_id = u.user_id
LEFT JOIN 
    follows ff ON ff.follower_id = @cur_user_id AND ff.following_id = u.user_id
LEFT JOIN repos r ON r.user_id = u.user_id AND (
        r.visibility_level = 'public' OR 
        r.visibility_level = 'signed' OR
        (r.visibility_level = 'chosen' AND EXISTS(SELECT 1 FROM repo_relations WHERE repo_relations.repo_id = r.repo_id AND repo_relations.user_id = @cur_user_id AND repo_relations.relation_type = 'visi')) OR
        ((r.visibility_level = 'private' OR r.visibility_level = 'chosen') AND (r.user_id = @cur_user_id OR @role::text='admin')))
WHERE 
    f.follower_id = @user_id and u.fts_username @@ plainto_tsquery(@query) AND (u.blocked='false' OR @role::text='admin');
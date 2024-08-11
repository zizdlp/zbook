-- name: CreateVerification :one
INSERT INTO verifications (
  verification_url,
  user_id,
  verification_type
) VALUES (
  $1, $2,$3
) RETURNING *;

-- name: GetVerification :one
SELECT verifications.*,users.username,users.email
FROM verifications
JOIN users  ON users.user_id = verifications.user_id
WHERE verification_url = $1 
    AND is_used = FALSE
    AND expired_at > now() 
    LIMIT 1
FOR NO KEY UPDATE;

-- name: MarkVerificationAsUsed :one
UPDATE verifications
SET
    is_used = TRUE
WHERE
    verification_url = @verification_url
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;
-- name: CreateInvitation :one
INSERT INTO invitations (
  email,
  invitation_url
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetInvitation :one
SELECT *
FROM invitations
WHERE email = $1 AND invitation_url = $2;

-- name: MarkInvitationAsUsed :one
UPDATE invitations
SET
    is_used = TRUE
WHERE
    email = $1 AND invitation_url = $2
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;
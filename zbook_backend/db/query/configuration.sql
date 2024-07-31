-- name: UpdateConfiguration :exec
UPDATE configurations
SET config_value=$2,updated_at=now()
WHERE config_name=$1;

-- name: GetConfiguration :one
SELECT *
FROM configurations
WHERE config_name=$1
LIMIT 1;
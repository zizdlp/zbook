-- name: UpdateConfiguration :exec
UPDATE configurations
SET config_value=$2
WHERE config_name=$1;

-- name: GetConfiguration :one
SELECT *
FROM configurations
WHERE config_name=$1
LIMIT 1;
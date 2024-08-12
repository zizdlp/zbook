-- name: GetGeoInfo :one
SELECT  *
FROM
    geoip
WHERE
    $1::inet << ip_range_cidr;

-- name: GetGeoInfoBatch :many
SELECT *
FROM geoip
WHERE ip_range_cidr && ANY($1::inet[]);
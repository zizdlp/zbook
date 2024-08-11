-- name: GetGeoInfo :one
SELECT  *
FROM
    geoip
WHERE
    $1::inet << ip_range_cidr;
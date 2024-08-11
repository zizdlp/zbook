python convert.py
# pg_dump -U root -d zbook -t geoip --data-only > geoip_data.sql # in database docker container
# gzip geoip_data.sql # in database docker container
# docker cp zbook-local-database:/geoip_data.sql.gz .
# psql -U root -d zbook -f geoip_data.sql

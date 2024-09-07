CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
POD_NAME := zbook-local-database
include zbook_backend/app.env
########################################################
################## local backend  ######################
tidy:
	cd zbook_backend && \
	go mod tidy  
md2html:
	cd zbook_backend && \
	go run cmd/md2html/main.go ${CURRENT_DIR}/zbook_data/convert_src ${CURRENT_DIR}/zbook_data/convert_dest
compress:
	cd zbook_backend && \
	go run cmd/compress/main.go ${CURRENT_DIR}/zbook_data/source.png  ${CURRENT_DIR}/zbook_data/dest.png
server:
	cd zbook_backend && \
	REQUIRE_EMAIL_VERIFY=false go run cmd/server/main.go
mail:
	cd zbook_backend && \
	go run cmd/mail/main.go
gp:
	cd zbook_backend && \
	mkdir -p pb && \
	rm -f -r statik/* && \
	rm -f -r pb/* &&  \
	rm -f -r doc/swagger/*  && \
	protoc  --proto_path=proto --proto_path=proto/models --proto_path=proto/rpcs --go_out=pb --go_opt=paths=source_relative  \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative  \
	--grpc-gateway_out=pb \
	--openapiv2_out=doc/swagger \
	--grpc-gateway_opt paths=source_relative \
	proto/**/*.proto proto/*.proto && \
	statik -src=./doc -dest=./
sqlc:
	cd zbook_backend && \
	sqlc generate
mock:
	cd zbook_backend && \
	mockgen -package mockdb -destination db/mock/store.go github.com/zizdlp/zbook/db/sqlc Store && \
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/zizdlp/zbook/worker TaskDistributor
test:
	cd zbook_backend && \
	go test -v -cover -short ./... 

#########################################################
################## local frontend  ######################
npm_install:
	cd zbook_frontend && \
	npm install
npm_build:
	cd zbook_frontend && \
	npm run build
jtest:
	cd zbook_frontend && \
	npm run test
next:
	cd zbook_frontend && \
	npm run dev
#########################################################
################## local database  ######################
database:
	docker run --name zbook-local-database -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d zizdlp/zbook_database
createdb:
	docker exec -it zbook-local-database createdb --username=root --owner=root zbook
dropdb:
	docker exec -it zbook-local-database dropdb zbook
migrateup:
	cd zbook_backend && \
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up
migratedown:
	cd zbook_backend && \
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down
mmdb2psql:
	cd zbook_backend && \
	go run cmd/mmdb2psql/main.go  ${CURRENT_DIR}/GeoLite2-City.mmdb
download_data:
	pg_dump -U root -d zbook -t geoip --data-only > geoip_data.sql # in database docker container
	gzip geoip_data.sql # in database docker container
	docker cp zbook-local-database:/geoip_data.sql.gz .
	psql -U root -d zbook -f geoip_data.sql
import_data:
	docker cp geoip_data.sql.gz $(POD_NAME):/tmp/
	docker exec -it $(POD_NAME) sh -c "gunzip /tmp/geoip_data.sql.gz"
	docker exec -it $(POD_NAME) sh -c 'psql -U root -d zbook -f /tmp/geoip_data.sql'
	docker exec -it $(POD_NAME) sh -c "rm /tmp/geoip_data.sql"
redis:
	docker run --name zbook-local-redis -p 6379:6379 -d redis:7-alpine
minio:
	docker run --name zbook-local-minio -itd -p 9000:9000 -p 9001:9001 -e "MINIO_ROOT_USER=${MINIO_ROOT_USER}" -e "MINIO_ROOT_PASSWORD=${MINIO_ROOT_PASSWORD}" minio/minio server /data --console-address ":9001"
create_bucket:
	mc alias set avatar http://localhost:9000 ${MINIO_ROOT_USER} ${MINIO_ROOT_PASSWORD}
	mc mb avatar/avatar
evans:
	evans --host localhost --port 9090 -r repl
#########################################################################
################## build docker images from source ######################
build_frontend:
	sudo docker buildx build -t zizdlp/zbook_frontend:local -f ./zbook_frontend/zbook_frontend.Dockerfile ./zbook_frontend
build_backend:
	sudo docker buildx build -t zizdlp/zbook_backend:local -f ./zbook_backend/zbook_backend.Dockerfile ./zbook_backend
build_database:
	sudo docker buildx build -t zizdlp/zbook_database:local -f ./zbook_database/zbook_database.Dockerfile ./zbook_database

#########################################################################
################## run as docker compose ################################
compose_build:
	docker-compose --env-file compose.env -f docker-compose-build.yaml down --volumes
	sudo docker-compose --env-file compose.env -f docker-compose-build.yaml build
	docker-compose --env-file compose.env -f docker-compose-build.yaml up  --remove-orphans 
compose_pull:
	docker-compose --env-file compose.env -f docker-compose-pull.yaml down --volumes
	docker-compose --env-file compose.env -f docker-compose-pull.yaml pull
	docker-compose --env-file compose.env -f docker-compose-pull.yaml up --remove-orphans
.PHONY: database createdb dropdb migrateup migratedown sqlc mock test server next compose

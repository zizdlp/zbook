include zbook_backend/app.env
export

CURRENT_DIR := $(shell pwd)


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
##########################################
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
jtest:
	cd zbook_frontend && \
	npm run test
redis:
	docker run --name zbook-local-redis -p 6379:6379 -d redis:7-alpine
minio:
	docker run --name zbook-local-minio -itd -p 9000:9000 -p 9001:9001 -e "MINIO_ROOT_USER=${MINIO_ROOT_USER}" -e "MINIO_ROOT_PASSWORD=${MINIO_ROOT_PASSWORD}" minio/minio server /data --console-address ":9001"
create_bucket:
	mc alias set avatar http://localhost:9000 ${MINIO_ROOT_USER} ${MINIO_ROOT_PASSWORD}
	mc mb avatar/avatar
server:
	cd zbook_backend && \
	REQUIRE_EMAIL_VERIFY=false go run cmd/server/main.go
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
evans:
	evans --host localhost --port 9090 -r repl
batch_test:
	cd zbook_backend && \
	go run cmd/batch_test/main.go
build_frontend_localhost:
	docker build -t zbook_frontend_localhost --build-arg ENV_FILE=.env.production.localhost -f ./zbook_frontend/zbook_frontend.Dockerfile ./zbook_frontend
build_frontend_zizdlp:
	docker build -t zbook_frontend_zizdlp --build-arg ENV_FILE=.env.production.zizdlp -f ./zbook_frontend/zbook_frontend.Dockerfile ./zbook_frontend
build_backend:
	docker build -t zbook_backend --build-arg BUILDPLATFORM=amd64 -f ./zbook_backend/zbook_backend.Dockerfile ./zbook_backend
build_database:
	docker build -t zbook_database -f ./zbook_database/zbook_database.Dockerfile ./zbook_database
run_frontend:
	docker run -it zizdlp/zbook_frontend

#frontend
next:
	cd zbook_frontend && \
	npm run dev
compose:
	sudo docker-compose -f docker-compose.yaml down --volumes
	sudo docker-compose -f docker-compose.yaml build
	sudo docker-compose -f docker-compose.yaml up  --remove-orphans 
.PHONY: database createdb dropdb migrateup migratedown sqlc mock test server next compose

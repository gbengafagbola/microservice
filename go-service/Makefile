DB_URL=postgresql://root:hush@localhost:5432/go_service_db?sslmode=disable
postgres:
	docker run --name go-service-postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=hush -d postgres:16-alpine
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)	
createdb:
	docker exec -it go-service-postgres-container createdb --username=root --owner=root go_service_db
dropdb:
	docker exec -it go-service-postgres-container dropdb go_service_db	
sqlc:
	sqlc generate	
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=go-microservice \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/gbengafagbola/microservice/go-service/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/gbengafagbola/microservice/go-service/worker TaskDistributor	
evans:
	evans --host localhost --port 9090 -r repl
redis:
	docker run --name redis -p 6379:6379 -d redis:7.2.4-alpine
server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown new_migration sqlc proto evans redis server
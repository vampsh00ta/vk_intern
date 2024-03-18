DATABASE="postgresql://vk_intern:vk_intern@localhost:5432/vk_intern?sslmode=disable"
migrate:
	migrate create -ext sql -dir ./migrations/ -seq $(name)
migration:
	migrate -path ./migrations -database  $(DATABASE)  up
start:
	docker-compose build &&  docker-compose up
http-tests:
	go test ./internal/transport/http/v1/tests
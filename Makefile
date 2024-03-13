DATABASE="postgresql://vk_intern:vk_intern@localhost:5432/vk_intern?sslmode=disable"
migrate:
	migrate create -ext sql -dir ./migrations/ -seq $(name)
migration:
	migrate -path ./migrations -database  $(DATABASE)  up


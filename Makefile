APP_NAME=users-api
BUILD_PATH=bin

.PHONY: clean
clean:
	rm -rf $(BUILD_PATH)/
	go mod tidy

.PHONY: build
build: clean
	go build -o $(BUILD_PATH)/$(APP_NAME) main.go

start:
	go run main.go

create-local-db:
	docker run --name users-db -d \
	-e POSTGRES_USER=users-api \
	-e POSTGRES_PASSWORD=aaa \
	-p 5432:5432 postgres:16-bookworm
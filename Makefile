APP_NAME=users-api
BUILD_PATH=bin

.PHONY: clean
clean:
	rm -rf $(BUILD_PATH)/
	go mod tidy

.PHONY: build
build: clean
	go build -o $(BUILD_PATH)/$(APP_NAME) main.go
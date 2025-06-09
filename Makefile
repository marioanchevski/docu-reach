APP_NAME = docu-reach

run: build
	@./bin/$(APP_NAME)

build: 
	@go build -o bin/$(APP_NAME) cmd/main.go

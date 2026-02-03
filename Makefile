default: build

APP_NAME := "xmlminer"
CMD_DIR := "cmd/$(APP_NAME)/main.go"

.PHONY: build
build:
	go build -o $(APP_NAME) $(CMD_DIR)

.PHONY: push
push:
	git add -A
	git commit -m "."
	git push

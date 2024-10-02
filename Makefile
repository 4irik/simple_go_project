SHELL = /bin/sh
DC_CONTAINER_NAME = golang:1.23.0-alpine
DC_USER = $(shell id -u)
DC_GROUP = $(shell id -g)
DC_RUN_ARGS_DEF = --rm -it -v ./:/go -w /go
DC_RUN_ARGS_USER = --user $(DC_USER):$(DC_GROUP)
DC_RUN_PREPARE = export HOME=/tmp/

help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  \033[32m%-18s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


root-shell: ## App container shell as root
	docker run $(DC_RUN_ARGS_DEF) $(DC_CONTAINER_NAME) $(SHELL)

shell: ## App container shell
	docker run $(DC_RUN_ARGS_DEF) $(DC_RUN_ARGS_USER) $(DC_CONTAINER_NAME) $(SHELL) -c "${DC_RUN_PREPARE} && ${SHELL}"

build: ## Build app
	docker run $(DC_RUN_ARGS_DEF) $(DC_RUN_ARGS_USER) $(DC_CONTAINER_NAME) $(SHELL) -c "${DC_RUN_PREPARE} && go build -v"

test: ## Run tests
	docker run $(DC_RUN_ARGS_DEF) $(DC_RUN_ARGS_USER) $(DC_CONTAINER_NAME) $(SHELL) -c "${DC_RUN_PREPARE} && go test ./..."

run: ## Run app
	docker run $(DC_RUN_ARGS_DEF) $(DC_RUN_ARGS_USER) $(DC_CONTAINER_NAME) $(SHELL) -c "${DC_RUN_PREPARE} && go run main.go"

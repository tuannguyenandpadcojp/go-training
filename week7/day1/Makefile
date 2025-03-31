help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


db-migration: ## Run DB migration
	@migrate -path database/migrations -database "mysql://root:rootpassword@tcp(localhost:3306)/tm?multiStatements=true" up

exec-db: ## Exec into the database
	@docker exec -it mysql mysql -uroot -prootpassword tm

up: ## Start the application dependencies
	@docker-compose up -d

down: ## Stop the application dependencies
	@docker-compose down

run:
	@export CONFIG_PATH=.local.env && go run cmd/main.go

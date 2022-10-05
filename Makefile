.PHONY: help openapi api-up api-down api-logs ps 
.DEFAULT_GOAL := help

openapi: ## generates codes by openapi-generator
	docker run --rm -v ${PWD}:/root -v ${PWD}/spec:/spec  openapitools/openapi-generator-cli:v6.0.0 generate -g go-server -i /spec/igusaya_blog.yaml -o /root/api/gen --additional-properties=packageName=openapi,router=chi,sourceFolder=openapi
	goimports -w api/gen/openapi/*

api-up: ## Do docker compose up with hot reload
	docker compose -f ./api/docker-compose.yml up -d

api-down: ## Do docker compose down
	docker compose -f ./api/docker-compose.yml down

api-logs: ## Tail docker compose logs
	docker compose -f ./api/docker-compose.yml logs -f

api-ps: ## Check container status
	docker compose -f ./api/docker-compose.yml ps

dry-migrate: ## Try migration
	mysqldef -u user -p pass -h 127.0.0.1 -P 33306 blog --dry-run < ./api/_tools/mysql/schema.sql

migrate:  ## Execute migration
	mysqldef -u user -p pass -h 127.0.0.1 -P 33306 blog < ./api/_tools/mysql/schema.sql

mysql: ## run mysql
	mysql -h 127.0.0.1 -P 33306 -u todo -ptodo

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
.PHONY: help openapi
.DEFAULT_GOAL := help

openapi: ## generates codes by openapi-generator
	docker run --rm -v ${PWD}:/root -v ${PWD}/spec:/spec  openapitools/openapi-generator-cli:v6.0.0 generate -g go-server -i /spec/igusaya_blog.yaml -o /root/api/gen --additional-properties=packageName=openapi,router=chi,sourceFolder=openapi
	goimports -w api/gen/openapi/*

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
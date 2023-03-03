include .env
BINARY_NAME=fyfirman-blog-service

## Build:
build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
	GO111MODULE=on go build -o out/bin/$(BINARY_NAME) ./cmd

watch: ## Watch your project and rebuild it on changes
	gin --appPort 8080 --path ./cmd run main.go

## Docker
docker-build: ## Build docker image
	docker build -t fyfirman-blog-service .

docker-run: ## Run docker image
	docker run -p 8080:8080 -e FIREBASE_DATABASE_URL=${FIREBASE_DATABASE_URL} -e SERVER_PORT=${SERVER_PORT} -v $(shell pwd)/serviceAccountKey.json:/app/serviceAccountKey.json fyfirman-blog-service 

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
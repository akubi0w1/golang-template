OAS_GEN_PACKAGE_NAME=oapi
OAS_PATH=./oas/openapi.yaml
OAS_GEN_OUT_DIR=./interface/oapi
MOCK_SERVER_OUT_DIR=./mock-server
ENT_PATH=./interface/database/ent
GO_MOCK_DIRS=./usecase ./domain/service ./domain/repository
BUILD_DIR=./binary
BUILD_NAME=api

help: ## display help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

gen-mock: ## generate code of gomock
	for package in $(GO_MOCK_DIRS); do \
		go generate -v $${package}; \
	done

build: ## build api
	go build -o $(BUILD_DIR)/$(BUILD_NAME) ./cmd/...

test: ## run go test
	go test ./...

dev-up: ## up develop env
	docker-compose up -d

dev-start: ## start develop env
	docker-compose start

dev-stop: ## stop develop env
	docker-compose stop

dev-down: ## down develop env
	docker-compose down --rmi local --volumes

dev-log-api: ## watch api's log
	docker-compose logs -f api

dev-log-db: ## watch db's log
	docker-compose logs -f db

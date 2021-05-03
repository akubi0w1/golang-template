OAS_PATH=./oas/openapi.yaml
MOCK_SERVER_OUT_DIR=./mock-server
ENT_PATH=./interface/persistent/mysql/ent
GO_MOCK_DIRS=./usecase ./domain/service ./domain/repository
BUILD_DIR=./build
BUILD_NAME=api

help: ## display help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

gen-schema: ## generate code of ent
	go generate $(ENT_PATH)

gen-mock: ## generate code of gomock
	for package in $(GO_MOCK_DIRS); do \
		go generate -v $${package}; \
	done

build: ## build api
	go build -o $(BUILD_DIR)/$(BUILD_NAME) ./cmd/...

test: ## run go test
	go test --cover ./...

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

doc-up: ## run swagger ui
	docker run -d -p 8888:8080 -v `pwd`:/usr/share/nginx/html/api -e API_URL=api/$(OAS_PATH) swaggerapi/swagger-ui

doc-down: ## down swagger ui

build-mock: ## build mock server
	rm -rf $(MOCK_SERVER_OUT_DIR)
	docker run --rm -v `pwd`:/work swaggerapi/swagger-codegen-cli-v3:3.0.25 generate \
		-l nodejs-server \
		-i /work/$(OAS_PATH) \
		-o /work/$(MOCK_SERVER_OUT_DIR)

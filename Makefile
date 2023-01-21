PROJECT_NAME=lobster
SVC_NAME=
SVC_ADAPTER=
APP_NAME=$(PROJECT_NAME)-$(SVC_NAME)-$(SVC_ADAPTER)
MAIN_FOLDER='./cmd/$(SVC_ADAPTER)/$(SVC_NAME)'
VERSION=latest
DEPLOY_TO=uat
NS=$(DEPLOY_TO)-$(PROJECT_NAME)
REGISTRY=
HELM_REPO_NAME=

target_list_of_package = package-$(PROJECT_NAME)
target_list_of_push_chart = push-chart-$(PROJECT_NAME)

.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## clean artifacts
	@rm -rf bin charts coverage.txt profile.out
	@echo Successfully removed artifacts

.PHONY: lint
lint: ## execute golint
	@golangci-lint run ./... -c .golangci.yaml

.PHONY: test-unit
test-unit: ## execute unit test
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: test-e2e
test-e2e: ## execute e2e test
	@cd ./test/e2e && npx playwright test ./tests

.PHONY: build-image
build-image: check-SVC_NAME check-SVC_ADAPTER check-VERSION check-GITHUB_TOKEN ## build docker image with APP_NAME and VERSION
	@docker build -t $(REGISTRY)/$(APP_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg MAIN_FOLDER=$(MAIN_FOLDER) \
	--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
	--platform linux/amd64 \
	--pull --cache-from=$(REGISTRY)/$(APP_NAME) \
	-f Dockerfile .

.PHONY: list-images
list-images: check-SVC_NAME check-SVC_ADAPTER ## list all images
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images: check-SVC_NAME check-SVC_ADAPTER ## remove all images
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image
push-image: check-SVC_NAME check-SVC_ADAPTER check-VERSION ## push image to registry
	@docker push $(REGISTRY)/$(APP_NAME):$(VERSION)

.PHONY: gen
gen: gen-pb gen-wire gen-mocks gen-swagger ## generate all generate commands

.PHONY: gen-wire
gen-wire: ## generate wire code
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -q --dir ./cmd/restful/account,./cmd/restful/rental,./ -o ./api/docs
	## Generated swagger spec

.PHONY: gen-pb
gen-pb: ## generate protobuf messages and services
	@go get -u google.golang.org/protobuf/proto
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=./pb --go_out=paths=source_relative:./pkg/entity --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./pb ./pb/domain/*/*/*.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./pkg/entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

.PHONY: gen-mocks
gen-mocks: ## generate mocks code via mockery
	@go generate -tags=wireinject -x ./...

.PHONY: package-all
package-all: $(target_list_of_package) ## package all helm charts

.PHONY: $(target_list_of_package)
$(target_list_of_package): ## package helm chart of application
	@helm package ./app/helm/$(subst package-,,$@) -d ./charts

.PHONY: push-chart-all
push-chart-all: $(target_list_of_push_chart) ## push all helm charts

.PHONY: $(target_list_of_push_chart)
$(target_list_of_push_chart): ## push helm chart of application
	@helm s3 push --force ./charts/$(subst push-chart-,,$@)-* $(HELM_REPO_NAME)
	@echo Successfully pushed $(subst push-chart-,,$@) chart

.PHONY: deploy
deploy: check-SVC_NAME check-SVC_ADAPTER check-VERSION check-DEPLOY_TO ## deploy the application via helm 3
	@helm -n $(NS) upgrade --install $(DEPLOY_TO)-$(APP_NAME) \
	$(HELM_REPO_NAME)/$(PROJECT_NAME) \
	--set "image.tag=$(VERSION)" -f $(MAIN_FOLDER)/configs/$(DEPLOY_TO).yaml

.PHONY: update-package
update-package: ## update package and commit
	@go get -u ./...
	@go mod tidy

	@bazel run //:gazelle-update-repos

	@git add go.mod go.sum deps.bzl
	@git commit -m "build: update package"

.PHONY: test-smoke
test-smoke: ## execute smoke testing
	@k6 run test/k6/requests/*

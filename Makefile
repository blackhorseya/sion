## env for project
PROJECT_NAME := irent
VERSION := $(shell git describe --tags --always)

## env for helm
HELM_REPO_NAME := sean-side
RELEASE_NAME := $(DEPLOY_TO)-$(PROJECT_NAME)
NS := $(PROJECT_NAME)
DEPLOY_TO := uat

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

.PHONY: gazelle-repos
gazelle-repos: ## update gazelle repos
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune

.PHONY: gazelle
gazelle: gazelle-repos ## run gazelle with bazel
	@bazel run //:gazelle

.PHONY: build-go
build-go: gazelle ## build go binary
	@bazel build //...

.PHONY: test-go
test-go: gazelle ## test go binary
	@bazel test //...

## docker
.PHONY: push-irent-app-restful-image
push-irent-app-restful-image: ## push irent app restful image
	@bazel run //:$@ --define=VERSION=$(VERSION)

.PHONY: push-irent-rental-cronjob-image
push-irent-rental-cronjob-image: ## push irent rental cronjob image
	@bazel run //:$@ --define=VERSION=$(VERSION)

.PHONY: gen
gen: gen-pb gen-wire gen-go gen-swagger ## generate all generate commands

.PHONY: gen-go
gen-go: ## generate go code
	@go generate -tags="wireinject" ./...

.PHONY: gen-wire
gen-wire: ## generate wire code
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -q --dir ./cmd/restful/app,./ -o ./api/docs
	## Generated swagger spec

.PHONY: gen-pb
gen-pb: ## generate protobuf messages and services
	@go get -u google.golang.org/protobuf/proto
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=. \
			--go_out=. --go_opt=module=github.com/blackhorseya/irent \
			--go-grpc_out=. --go-grpc_opt=module=github.com/blackhorseya/irent,require_unimplemented_servers=false \
			./pb/domain/*/**.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./pkg/entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

.PHONY: deploy
deploy: check-SVC_NAME check-SVC_ADAPTER check-VERSION check-DEPLOY_TO ## deploy the application via helm 3
	@helm -n $(NS) upgrade --install $(DEPLOY_TO)-$(APP_NAME) \
	$(HELM_REPO_NAME)/$(PROJECT_NAME) \
	--set "image.tag=$(VERSION)" -f $(MAIN_FOLDER)/configs/$(DEPLOY_TO).yaml

.PHONY: update-package
update-package: ## update package and commit
	@go get -u ./...
	@go mod tidy

	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune

	@git add go.mod go.sum deps.bzl
	@git commit -m "build: update package"

.PHONY: test-smoke
test-smoke: ## execute smoke testing
	@k6 run test/k6/requests/*

N = 1
DB_URI='mysql://root:changeme@tcp(localhost:3306)/$(PROJECT_NAME)_$(SVC_NAME)?charset=utf8mb4&parseTime=True&loc=Local'

.PHONY: migrate-up
migrate-up: check-SVC_NAME ## run migration up
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migration/$(SVC_NAME) up

.PHONY: migrate-down
migrate-down: check-SVC_NAME check-N ## run migration down
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migration/$(SVC_NAME) down $(N)

## helm
.PHONY: lint-helm
lint-helm: ## lint helm chart
	@helm lint deployments/charts/*

.PHONY: add-helm-repo
add-helm-repo: ## add helm repo
	@helm repo add --no-update $(HELM_REPO_NAME) gs://sean-helm-charts/charts
	@helm repo update $(HELM_REPO_NAME)

.PHONY: package-helm
package-helm: ## package helm chart
	@helm package deployments/charts/* --destination deployments/charts

.PHONY: push-helm
push-helm: ## push helm chart
	@helm gcs push --force ./deployments/charts/*.tgz $(HELM_REPO_NAME)
	@helm repo update $(HELM_REPO_NAME)

.PHONY: upgrade-helm
upgrade-helm: ## upgrade helm chart
	@echo "Upgrading $(RELEASE_NAME) to $(VERSION)"
	@echo "Using config: ./deployments/configs/$(DEPLOY_TO)/values.yaml"
	@helm upgrade $(RELEASE_NAME) $(HELM_REPO_NAME)/$(PROJECT_NAME) \
	--install --namespace $(NS) \
	--history-max 3 \
	--values ./deployments/configs/$(DEPLOY_TO)/values.yaml \
	--set version=$(VERSION)

## deployments
INCREMENT := PATCH

.PHONY: release
release: ## release this application
	@cz bump --changelog --yes -s --increment=$(INCREMENT)
	@git push && git push --tags
	@echo "Version: $(VERSION) to $(DEPLOY_TO)"

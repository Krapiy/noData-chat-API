GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
PROJECT_FILES=$(go list ./... | grep -v /vendor/)
BINARY_NAME=noData-chat-API
COVER_FILE=coverage.out

## Develop
run_development: build_development start_development

build_development:
	$(GOGET) github.com/codegangsta/gin
	$(GOGET) ./...

start_development:
	gin --appPort=${PORT} --port=${DEV_PORT} run main.go

## Common
run_test:
	$(GOTEST) ./... -v -count=1

test_coverage:
	$(GOTEST) ./... -coverprofile=$(COVER_FILE) && $(GOTOOL) cover -html=$(COVER_FILE)

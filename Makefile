# Go parameters
GOENV=CGO_ENABLED=0 GOFLAGS="-count=1"
GOCMD=$(GOENV) go
GOGET=go get -u
GOTEST=$(GOCMD) test -covermode=atomic -coverprofile=./coverage.txt
LOCAL_DB=localhost
DATETIME:= $(shell /bin/date "+%Y%m%d%H%M%S")
DB_HOST=localhost
DB_NAME=test
DB_USER=test
DB_PASSWORD=test
DB_SSLMODE=disable

install:
	GO111MODUlE=on $(GOCMD) mod vendor

# Use this to run all tests
tests:
	$(GOTEST) ./...

# Use this to run a single test.
# Example usage: make test test=TestPing
test:
	$(GOTEST) ./... -ginkgo.focus=${test}

run:
	$(GOCMD) run main.go

local-tests: pg-stop pg
	$(GOTEST) ./...

local-test: pg-stop pg
	$(GOTEST) ./... -ginkgo.focus=${test}

pg-stop:
	@echo "[`date`] Stopping previous launched postgres [if any]"
	docker stop pg || true

pg:
	@echo "[`date`] Starting Postgres container"
	docker run -d --rm --name pg \
    -p 5432:5432 \
    -e POSTGRES_DB=${DB_NAME} \
    -e POSTGRES_USER=${DB_USER} \
    -e POSTGRES_PASSWORD=${DB_PASSWORD} \
    postgres:9.6
	sleep 3

local-run: pg-stop pg
	$(GOCMD) run main.go

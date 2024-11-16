PHONY: run
run:
	#go run ./cmd/main.go dev -v $(PARAMS)
	go run ./cmd/main.go doctor

build-linux64:
	env GOOS=linux GOARCH=amd64 go build -o ./dist/bootcamp-cli ./cmd/main.go

# Makefile for running tests in different parts of the project

# Variables
GO := go
TEST_DIR := ./test
DOMAIN_DIR := ./domain
APPLICATION_DIR := ./application
E2E_DIR := ./test/e2e
CMD_DIR := ./cmd

# Default target to run all tests
.PHONY: all
all: domain_tests application_tests e2e_tests

# Target for Domain tests
.PHONY: domain_tests
domain_tests:
	$(GO) test $(DOMAIN_DIR) -v

# Target for Application tests
.PHONY: application_tests
application_tests:
	$(GO) test $(APPLICATION_DIR) -v

# Target for E2E tests
.PHONY: e2e_tests
e2e_tests:
	$(GO) test $(E2E_DIR) -v

# Target to run CLI tests (if needed for the cmd directory)
.PHONY: cli_tests
cli_tests:
	$(GO) test $(CMD_DIR) -v

# Target to run all tests in the project (Domain, Application, E2E)
.PHONY: run_all_tests
run_all_tests:
	$(MAKE) domain_tests
	$(MAKE) application_tests
	$(MAKE) e2e_tests

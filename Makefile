help:
	@echo "Usage: 'make <target>' where <target> is one of:"
	@echo "  ci       to run CI tests."
	@echo "  coverage to generate coverage report."
	@echo "  lint     to check code linting."

ci:
	go test -race -mod=vendor ./...

cover:
	go test -coverprofile=coverage.txt ./... -mod=vendor

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint

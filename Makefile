mod install: export GO111MODULE=on
mod install: export GOPROXY=direct
mod install: export GOSUMDB=off

.PHONY: mod
## Install project dependencies using go mod. Usage 'make mod'
mod:
	@go mod tidy
	@go mod vendor

.PHONY: run
## Run service. Usage: 'make run'
run: ; $(info Starting svc...)
	go run --tags dev ./cmd/server/.

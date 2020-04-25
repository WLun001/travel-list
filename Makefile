# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m
# Define header
HEADER=$(GREEN)Fiber Recipe$(NOCOLOR)

AIR := $(shell command -v air || 0)
.PHONY: help
default: help

help: ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

.PHONY: install
install: ## install dependencies
	@make -j2 install-web install-go
	@make tool
.PHONY: start
start: ## start web and Go with live reload
	@make -j2 start-web start-go

.PHONY: start-go
start-go: ## start Go with lie reload
ifdef AIR
	@echo "$(HEADER): Starting go"
	@AIR
else
	@echo "please intall air, run make tool"
endif

.PHONY: build
build: ## build docker image
	docker build -t travel:${TAG} .

.PHONY: install-go
install-go: ## download Go dependencies
	go mod download

.PHONY: clean-go
clean-go: ## remove go binary
	rm -rf tmp/
	go clean -modcache
	go mod tidy

.PHONY: tool
tool:
	go get -u github.com/cosmtrek/air

.PHONY: start-web
start-web: ## start web with live reload
	@echo "$(HEADER): Starting web"
	@npm start --prefix web

.PHONY: clean-web
clean-web: ## remove web/dist
	rm -rf web/dist

.PHONY: install-web
install-web: ## download web dependencies
	npm i --prefix web



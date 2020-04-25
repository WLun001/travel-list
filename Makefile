# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m
# Define header
HEADER=$(GREEN)Fiber Recipe$(NOCOLOR)

AIR := $(shell command -v air || 0)
.PHONY: help start start-go start-web clean-web clean-go build

default: help

help: ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

start: ## start web and Go with live reload
	@make -j2 start-web start-go

start-go: ## start Go with lie reload
ifdef AIR
	@echo "$(HEADER): Starting go"
	@AIR
else
	@echo "please intall air, run make tool"
endif

build: ## build docker image
	docker build -t travel:${TAG} .

download-go: ## download Go dependencies
	go mod download

clean-go: ## remove go binary
	rm -rf tmp/
	go clean -modcache
	go mod tidy

tool:
	go get -u github.com/cosmtrek/air

start-web: ## start web with live reload
	@echo "$(HEADER): Starting web"
	@npm start --prefix web

clean-web: ## remove web/dist
	rm -rf web/dist

download-web: ## download web dependencies
	npm i --prefix web



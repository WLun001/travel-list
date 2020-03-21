# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m
# Define header
HEADER=$(GREEN)Fiber Recipe$(NOCOLOR)

.PHONY: help start start-go start-web clean-web clean-go

default: help

help: ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

start-web: ## start web with live reload
	@echo "$(HEADER): Starting web"
	@npm run build --prefix web

start-go: ## start Go with lie reload
	@echo "$(HEADER): Starting go"
	@air

start: ## start web and Go with live reload
	@make -j2 start-web start-go

clean-web: ## remove web/dist
	rm -rf web/dist

clean-go: ## remove go binary
	rm -rf tmp/


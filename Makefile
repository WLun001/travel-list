# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m
# Define header
HEADER=$(GREEN)Fiber Recipe$(NOCOLOR)

.PHONY: start-go start-web

start-web:
	@echo "$(HEADER): Starting web"
	@npm run build --prefix web

start-go:
	@echo "$(HEADER): Starting go"
	@air


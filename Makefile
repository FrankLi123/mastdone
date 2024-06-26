.PHONY: format lint

format:
	@gofmt -w main/*
	@goimports -w main/*

lint:
	@golint main/*
.PHONY: elixir

elixir:
	go run main.go elixir

tests:
	go test -v ./... | { grep -v 'no test files'; true; }

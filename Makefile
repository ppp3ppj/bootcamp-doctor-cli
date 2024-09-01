PHONY: run
run:
	#go run ./cmd/main.go dev -v $(PARAMS)
	go run ./cmd/main.go doctor

build-linux64:
	env GOOS=linux GOARCH=amd64 go build -o ./dist/bootcamp-cli ./cmd/main.go

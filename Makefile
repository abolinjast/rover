install:
	go build -o ./bin/rover ./cmd/main.go

run: install
	./bin/rover

run dev:
	go run ./cmd/main.go
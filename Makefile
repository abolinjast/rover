install:
	go build -o ./bin/rover main.go

run: install
	./bin/rover

run dev:
	go run main.go
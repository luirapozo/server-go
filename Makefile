build:
	go build -o bin/client src/client.go
	go build -o bin/server src/server.go

runc:
	go run src/client.go

runs:
	go run src/server.go

.PHONY: clean
clean:
	mkdir -p bin/
	rm -f bin/client bin/server
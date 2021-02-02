SOURCES := $(shell find . -name '*.go')

start: db dev

kroki: $(SOURCES) go.mod go.sum
	go build -o $@ ./cmd/kroki

db:
	docker-compose up --detach --remove-orphans

dev: cert.pem
	gow -c -v run ./cmd/kroki

lint:
	golangci-lint run --tests=false ./...

cert.pem key.pem:
	go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" --host localhost

clean:
	rm -f cert.pem key.pem kroki

.PHONY: db dev lint clean

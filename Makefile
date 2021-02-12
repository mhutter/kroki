IMAGE := quay.io/mhutter/kroki:latest
SOURCES := $(shell find . -name '*.go')

start: dev

kroki: $(SOURCES) go.mod go.sum
	go build -o $@ ./cmd/kroki

db:
	docker-compose up --detach --remove-orphans

dev: cert.pem
	gow -c -v run ./cmd/kroki

lint:
	golangci-lint run --tests=false ./...

image:
	docker build -t $(IMAGE) .
push:
	docker push $(IMAGE)

cert.pem key.pem:
	go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" --host localhost

clean:
	rm -f cert.pem key.pem kroki

.PHONY: db dev lint clean image push

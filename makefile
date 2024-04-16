.PHONY:df service_up run

image_up:
	docker build --tag banner .

service_up:
	docker-compose -f docker-compose.yml up -d --remove-orphans

run:
	go run ./cmd/main.go

lint:
	golangci-lint run ./... --config=./config/.golangci.yml --fast
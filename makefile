.PHONY:lint df service_up run

lint:
	golangci-lint run ./... --config=./config/.golangci.yml --fast

df:
	docker build --tag banner .

service_up:
	docker-compose -f docker-compose.yml up -d --remove-orphans

run:
	go run ./cmd/main.go
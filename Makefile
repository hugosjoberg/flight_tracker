start:
	docker-compose up --build

lint:
	golangci-lint run
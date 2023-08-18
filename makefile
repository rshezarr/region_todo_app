run:
	go run ./cmd/app/main.go

run-db:
	docker run -d -p 27017:27017 --name=lists_db mongo

run-compose:
	docker-compose up --build

down-compose:
	docker-compose down
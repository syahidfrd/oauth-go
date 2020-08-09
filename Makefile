dev:
	go run main.go

db-docker:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres
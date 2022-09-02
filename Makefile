up:
	@echo "docker compose up"
	docker compose up
	@echo "running port on 8080"
	go run main.go

run:
	@echo "running port on 8080"
	go run main.go

dup:
	@echo "docker compose up"
	docker compose up		
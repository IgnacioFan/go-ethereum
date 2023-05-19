-include .env
export

dev.build:
	docker-compose -f docker-compose.dev.yml up --build

dev.start:
	docker-compose -f docker-compose.dev.yml up

dev.stop:
	docker compose down

app.build:
	docker-compose up --build

app.start:
	docker-compose up

db.cli:
	docker exec -it $(POSTGRES_HOST) psql -U $(POSTGRES_USER)

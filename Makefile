build:
	docker compose build

run:
	docker compose up

run-d:
	docker compose up -d

recreate:
	docker compose down
	docker compose up --build --force-recreate
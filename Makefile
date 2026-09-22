build:
	docker compose build

run:
	docker compose up

recreate:
	docker compose down
	docker compose up --build --force-recreate
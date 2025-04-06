build:
	sudo rm -rf ./infra/containers/development/db/data/*
	docker compose up -d --build
	make logs
run:
	docker compose up -d
	make logs
logs:
	docker compose logs -f api
restart:
	docker compose restart api
	make logs
stop:
	docker compose stop
create-docs:
	~/go/bin/swag init -g ./cmd/main.go -o cmd/docs

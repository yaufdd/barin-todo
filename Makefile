.PHONY: build up down logs test lint fmt migrate


# build docker images
build:
	docker compose build

# run all services in background
up:
	docker compose up -d

# stop all services
down:
	docker compose down

# clean all (containers, images, volumes)
clean:
	docker compose down -v --rmi all

# logs server
logs:
	docker compose logs -f server


# enter to mysql
mysql:
	docker compose exec db mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE

# apply migrations 
migrate-up:
	docker run --rm -v $$PWD/migration:/migrations --network host migrate/migrate -path=/migrations -database "mysql://$$MYSQL_USER:$$MYSQL_PASSWORD@tcp(localhost:3306)/$$MYSQL_DATABASE" up

# apply migrations down
migrate-down:
	docker run --rm -v $$PWD/migration:/migrations --network host migrate/migrate -path=/migrations -database "mysql://$$MYSQL_USER:$$MYSQL_PASSWORD@tcp(localhost:3306)/$$MYSQL_DATABASE" down


# CREATE TODO
client-create:
	docker compose run --rm client create --title "$(title)" --desc "$(desc)" --deadline "$(deadline)"

# UPDATE TODO
client-update:
	docker compose run --rm client update --id $(id) --title "$(title)" --desc "$(desc)" --status $(status)

# DELETE TODO
client-delete:
	docker compose run --rm client delete --id $(id)

# SHOW TODO LIST
client-list:
	docker compose run --rm client list --status $(status) --sort-by $(sortby) --sort-order $(sortorder)

client-list-all:
	docker compose run --rm client list


test:
	go test ./internal/repo



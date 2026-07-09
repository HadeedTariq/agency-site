DB_PROTOCOL=sqlite
DB_PATH=./db.sqlite3
MIGRATION_DIR=internal/db/migrations

migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)
up:
	./migrate.sh -p $(DB_PROTOCOL) -u $(DB_PATH)
down:
	./migrate.sh -p $(DB_PROTOCOL) -u $(DB_PATH) -d down
dev:
	air
MIGRATION_DIR=./migrations
DB_URL=postgres://postgres:password@localhost:5432/your_db?sslmode=disable

# Название новой миграции передаётся через `make create-migration name=init_course`
create-migration:
	@read -p "Enter name of migration: " migration \
	&& goose -dir $(MIGRATION_DIR) create $$migration sql

# Применить все миграции
migrate-up:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" up

# Откатить последнюю миграцию
migrate-down:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" down

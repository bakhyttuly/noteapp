DB_URL=postgres://postgres:akzharkyn@localhost:5432/noteapp?sslmode=disable
MIGRATIONS_PATH=db/migrations

migrate-up:
	.\migrate.exe -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	.\migrate.exe -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-version:
	.\migrate.exe -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

migrate-force:
	.\migrate.exe -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force 1

migrate-create:
	.\migrate.exe create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)
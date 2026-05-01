MIGRATE = C:/Users/User/go/bin/migrate.exe

DB_URL = postgres://postgres:123123@localhost:5432/digital_product_store?sslmode=disable

migrate-up:
	$(MIGRATE) -path migrations -database "$(DB_URL)" up

migrate-down:
	$(MIGRATE) -path migrations -database "$(DB_URL)" down 1

migrate-version:
	$(MIGRATE) -path migrations -database "$(DB_URL)" version

migrate-force:
	$(MIGRATE) -path migrations -database "$(DB_URL)" force $(version)

migrate-create:
	$(MIGRATE) create -ext sql -dir migrations $(name)
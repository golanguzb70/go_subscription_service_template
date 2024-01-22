run:
	go run cmd/main.go

migrate_up:
	migrate -path migrations/ -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable up

migrate_down:
	migrate -path migrations/ -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable down

migrate_force:
	migrate -path migrations/ -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable force 10

proto-gen:
	bash scripts/gen-proto.sh
run:
	@go run cmd/main/main.go || true

up:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/account_microservice?sslmode=disable' up

down:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/account_microservice?sslmode=disable' down
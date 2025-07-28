run:
	@go run cmd/main/main.go || true

up:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/pulse?sslmode=disable' up

down:
	@migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/pulse?sslmode=disable' down
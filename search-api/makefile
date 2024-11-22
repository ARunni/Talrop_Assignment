run:
	go run cmd/server/main.go

test:
	go test ./... -v
dep:
	go mod tidy
fmt:
	go fmt cmd/server/main.go
	go fmt config/config.go
	go fmt di/di.go
	go fmt internal/domain/user.go
	go fmt internal/handler/search_handler.go
	go fmt internal/infrastructure/database/postgres.go
	go fmt internal/repository/user_repository.go
	go fmt internal/usecase/search_usecase.go
	go fmt internal/utils/helper.go
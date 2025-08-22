.PHONY: help db-up db-down rest-server grpc-server test-rest test-grpc proto clean

# 預設目標
help:
	@echo "Available commands:"
	@echo "  make db-up        - Start MySQL database"
	@echo "  make db-down      - Stop MySQL database"
	@echo "  make rest-server  - Start the REST server"
	@echo "  make grpc-server  - Start the gRPC server"
	@echo "  make test-rest    - Test REST API endpoints"
	@echo "  make test-grpc    - Test gRPC endpoints"
	@echo "  make proto        - Generate protobuf Go files"
	@echo "  make clean        - Stop database and clean up"

# 啟動資料庫
db-up:
	@echo "Starting MySQL database..."
	docker-compose up -d mysql
	@echo "Waiting for database to be ready..."
	@sleep 10
	@echo "✅ Database is ready!"

# 停止資料庫
db-down:
	@echo "Stopping MySQL database..."
	docker-compose down
	@echo "✅ Database stopped!"

# 啟動 REST server
rest-server:
	@echo "Starting REST server..."
	go run cmd/rest-server/main.go

# 啟動 gRPC server
grpc-server:
	@echo "Starting gRPC server..."
	go run cmd/grpc-server/main.go

# 生成 protobuf Go 檔案
proto:
	@echo "Generating protobuf Go files..."
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/pkg/proto/book.proto
	@echo "✅ Protobuf files generated!"

# 測試 REST API
test-rest:
	@echo "Testing REST API endpoints..."
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh

# 測試 gRPC API
test-grpc:
	@echo "Testing gRPC endpoints..."
	@chmod +x scripts/test-grpc.sh
	@./scripts/test-grpc.sh

# 清理
clean: db-down
	@echo "Cleaning up..."
	@docker system prune -f
	@echo "✅ Cleanup completed!"

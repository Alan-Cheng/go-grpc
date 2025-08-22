.PHONY: help db-up db-down server test clean

# 預設目標
help:
	@echo "Available commands:"
	@echo "  make db-up     - Start MySQL database"
	@echo "  make db-down   - Stop MySQL database"
	@echo "  make server    - Start the REST server"
	@echo "  make test      - Test API endpoints"
	@echo "  make clean     - Stop database and clean up"

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

# 啟動 server
server:
	@echo "Starting REST server..."
	go run cmd/server/main.go

# 測試 API
test:
	@echo "Testing API endpoints..."
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh

# 清理
clean: db-down
	@echo "Cleaning up..."
	@docker system prune -f
	@echo "✅ Cleanup completed!"

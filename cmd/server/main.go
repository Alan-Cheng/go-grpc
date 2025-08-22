package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	restbooksserver "go-grpc/internal/apps/rest-books-server"
	"go-grpc/internal/pkg/repo"
	"go-grpc/internal/pkg/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// 資料庫連線設定
	dsn := "bookuser:bookpass@tcp(localhost:3306)/bookdb?charset=utf8mb4&parseTime=True&loc=Local"

	// 建立資料庫連線
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 建立 repository
	bookRepo := repo.GetNewBookRepository(db)

	// 建立 service
	bookService := service.GetNewBookService(bookRepo)

	// 建立 router
	router := restbooksserver.ProvideRouter(*bookService)

	// 設定 port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 啟動 server
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

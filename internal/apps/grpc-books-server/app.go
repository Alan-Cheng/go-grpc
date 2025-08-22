package grpcbooksserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go-grpc/internal/pkg/model"
	"go-grpc/internal/pkg/proto"
	"go-grpc/internal/pkg/repo"
	"go-grpc/internal/pkg/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	grpcServer *grpc.Server
	lis        net.Listener
}

type grpcBookServer struct {
	proto.UnimplementedBookServiceServer
	bookService *service.BookService
}

func (s *grpcBookServer) AddBook(ctx context.Context, req *proto.AddBookRequest) (*proto.AddBookResponse, error) {
	book := &model.DBBook{
		Isbn:      int(req.Book.Isbn),
		Name:      req.Book.Name,
		Publisher: req.Book.Publisher,
	}

	err := s.bookService.AddBook(book)
	if err != nil {
		return nil, err
	}

	return &proto.AddBookResponse{
		Message: "Book added successfully",
	}, nil
}

func (s *grpcBookServer) GetBook(ctx context.Context, req *proto.GetBookRequest) (*proto.GetBookResponse, error) {
	book, err := s.bookService.GetBook(int(req.Isbn))
	if err != nil {
		return nil, err
	}

	return &proto.GetBookResponse{
		Book: &proto.Book{
			Isbn:      int32(book.Isbn),
			Name:      book.Name,
			Publisher: book.Publisher,
		},
	}, nil
}

func (s *grpcBookServer) GetAllBooks(ctx context.Context, req *proto.Empty) (*proto.ListBookResponse, error) {
	books, err := s.bookService.GetAllBooks()
	if err != nil {
		return nil, err
	}

	var protoBooks []*proto.Book
	for _, book := range books {
		protoBooks = append(protoBooks, &proto.Book{
			Isbn:      int32(book.Isbn),
			Name:      book.Name,
			Publisher: book.Publisher,
		})
	}

	return &proto.ListBookResponse{
		Books: protoBooks,
	}, nil
}

func (s *grpcBookServer) UpdateBook(ctx context.Context, req *proto.UpdateBookRequest) (*proto.UpdateBookResponse, error) {
	book := &model.DBBook{
		Isbn:      int(req.Book.Isbn),
		Name:      req.Book.Name,
		Publisher: req.Book.Publisher,
	}

	err := s.bookService.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateBookResponse{
		Message: "Book updated successfully",
	}, nil
}

func (s *grpcBookServer) RemoveBook(ctx context.Context, req *proto.RemoveBookRequest) (*proto.RemoveBookResponse, error) {
	err := s.bookService.RemoveBook(int(req.Isbn))
	if err != nil {
		return nil, err
	}

	return &proto.RemoveBookResponse{
		Message: "Book removed successfully",
	}, nil
}

func NewApp() *App {
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

	// 建立 gRPC server
	grpcServer := grpc.NewServer()
	proto.RegisterBookServiceServer(grpcServer, &grpcBookServer{
		bookService: bookService,
	})

	// 啟用 reflection
	reflection.Register(grpcServer)

	// 設定 port
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "9090"
	}

	// 監聽 port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	return &App{
		grpcServer: grpcServer,
		lis:        lis,
	}
}

func (app *App) Start() {
	fmt.Printf("gRPC Server starting on port %s...\n", app.lis.Addr().String())

	// 優雅關閉
	go func() {
		if err := app.grpcServer.Serve(app.lis); err != nil {
			log.Fatal("Failed to serve:", err)
		}
	}()

	// 等待中斷信號
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func (app *App) Shutdown() {
	fmt.Println("Shutting down gRPC server...")
	app.grpcServer.GracefulStop()
	fmt.Println("gRPC server stopped")
}

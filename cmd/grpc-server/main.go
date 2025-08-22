package main

import (
	grpcbooksserver "go-grpc/internal/apps/grpc-books-server"
)

func main() {
	app := grpcbooksserver.NewApp()
	app.Start()
	app.Shutdown()
}

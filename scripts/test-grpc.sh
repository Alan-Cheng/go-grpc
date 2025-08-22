#!/bin/bash

echo "Testing gRPC endpoints..."

# 需要先安裝 grpcurl
# brew install grpcurl

GRPC_URL="localhost:9090"

echo "1. Getting all books..."
grpcurl -plaintext $GRPC_URL prot.BookService/GetAllBooks

echo -e "\n2. Getting book with ISBN 0..."
grpcurl -plaintext -d '{"isbn": 0}' $GRPC_URL prot.BookService/GetBook

echo -e "\n3. Adding a new book..."
grpcurl -plaintext -d '{"book": {"isbn": 5, "name": "Test Book", "publisher": "Test Publisher"}}' $GRPC_URL prot.BookService/AddBook

echo -e "\n4. Updating book with ISBN 0..."
grpcurl -plaintext -d '{"book": {"isbn": 0, "name": "Updated C Programming Language", "publisher": "Updated Prentice Hall"}}' $GRPC_URL prot.BookService/UpdateBook

echo -e "\n5. Getting all books again..."
grpcurl -plaintext $GRPC_URL prot.BookService/GetAllBooks

echo -e "\n6. Deleting book with ISBN 5..."
grpcurl -plaintext -d '{"isbn": 5}' $GRPC_URL prot.BookService/RemoveBook

echo -e "\n7. Final book list..."
grpcurl -plaintext $GRPC_URL prot.BookService/GetAllBooks

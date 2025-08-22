#!/bin/bash

echo "Testing Book API endpoints..."

BASE_URL="http://localhost:8080"

echo "1. Getting all books..."
curl -X GET "$BASE_URL/books" | jq '.'

echo -e "\n2. Getting book with ISBN 0..."
curl -X GET "$BASE_URL/books/0" | jq '.'

echo -e "\n3. Adding a new book..."
curl -X PUT "$BASE_URL/books" \
  -H "Content-Type: application/json" \
  -d '{"isbn": 5, "name": "Test Book", "publisher": "Test Publisher"}' | jq '.'

echo -e "\n4. Updating book with ISBN 0..."
curl -X POST "$BASE_URL/books" \
  -H "Content-Type: application/json" \
  -d '{"isbn": 0, "name": "Updated C Programming Language", "publisher": "Updated Prentice Hall"}' | jq '.'

echo -e "\n5. Getting all books again..."
curl -X GET "$BASE_URL/books" | jq '.'

echo -e "\n6. Deleting book with ISBN 5..."
curl -X DELETE "$BASE_URL/books/5" | jq '.'

echo -e "\n7. Final book list..."
curl -X GET "$BASE_URL/books" | jq '.'

package main

//go:generate go run main.go write_docs
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f ./docs/rula.json -A rula
//go:generate go run github.com/golang/mock/mockgen -source=./client/operations/operations_client.go -destination=./client/operations/mocks/mock_operations_client.go
import (
	_ "github.com/cucumber/godog/cmd/godog"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/swaggo/swag/cmd/swag"
)

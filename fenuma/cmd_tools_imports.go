package main

//go:generate go run main.go write_docs
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f ./docs/FENUMA.json -A fenuma
import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/swaggo/swag/cmd/swag"
)

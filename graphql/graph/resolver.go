package graph

import "gorm.io/gorm"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run generate.go
type Resolver struct {
	DB *gorm.DB
}

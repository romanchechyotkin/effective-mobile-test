package main

import (
	"github.com/romanchechyotkin/effective-mobile-test-task/internal/app"
)

// @title Swagger Documentation
// @version 1.0
// @description Effective Mobile test task in Gin Framework
// @host localhost:8080
func main() {
	app.NewApp().Run()
}

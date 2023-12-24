package main

import (
	_ "task-management/docs"
	"task-management/pkg/handlers/rest"
)

// @title Task Management API
// @version 1.0
// @description This is a sample Echo API with Swagger documentation.
// @host localhost:8080
// @BasePath /task-management/api/v1
func main() {
	e := rest.InitRouter()
	e.Logger.Fatal(e.Start("localhost:8080"))
}

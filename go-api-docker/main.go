package main

import (
	"go-api-docker/database"
	"go-api-docker/routes"
)

func init() {
	database.GetDBClient()
}

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}

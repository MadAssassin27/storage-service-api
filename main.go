package main

import (
	"storage-service-api/database"
	"storage-service-api/database/migration"
	"storage-service-api/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8000")
}

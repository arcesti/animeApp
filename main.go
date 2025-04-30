package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/arcesti/animeApp/routes"
)

func main() {
	app := fiber.New();

	routes.JinkanRoutes(app);

	log.Fatal(app.Listen(":3000"));
}
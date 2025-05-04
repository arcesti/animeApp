package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/arcesti/animeApp/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New();

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.Routes(app);

	log.Fatal(app.Listen(":3001"));
}
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/arcesti/animeApp/serviceApi"
)

func Routes (app *fiber.App) {

	animesRoute := app.Group("/anime");
	animesRoute.Get("/topPopular", serviceapi.GetTopAnimePopular());
	animesRoute.Get("/topAiring", serviceapi.GetTopAnimeAiring());
	animesRoute.Get("/:name", serviceapi.GetAnime());

	personagensRoute := app.Group("/personagens");
	personagensRoute.Get("/:id", serviceapi.GetPersonagens());
}


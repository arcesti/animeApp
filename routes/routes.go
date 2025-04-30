package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/arcesti/animeApp/model"
	"fmt"
	"net/http"
	"encoding/json"
)

type JikanResponse struct {
	Data []model.Anime `json:"data"`
}

func JinkanRoutes (app *fiber.App) {
	animeRoutes := app.Group("/anime");

	animeRoutes.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		apiURL := fmt.Sprintf("https://api.jikan.moe/v4/anime?q=%s&limit=1", name)

		resp, err := http.Get(apiURL)
		if err != nil {
			return c.Status(500).SendString("Erro ao buscar anime")
		}
		defer resp.Body.Close()

		var jikanResp JikanResponse
		if err := json.NewDecoder(resp.Body).Decode(&jikanResp); err != nil {
			return c.Status(500).SendString("Erro ao decodificar resposta")
		}

		if len(jikanResp.Data) == 0 {
			return c.Status(404).SendString("Anime não encontrado")
		}

		return c.JSON(jikanResp.Data[0])
	})
}
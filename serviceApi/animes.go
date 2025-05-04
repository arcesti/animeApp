package serviceapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arcesti/animeApp/model"
	"github.com/gofiber/fiber/v2"
)

type JikanResponse struct {
	Data []model.Anime `json:"data"`
}

func GetAnime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		apiURL := fmt.Sprintf("https://api.jikan.moe/v4/anime?q=%s&limit=1", name)

		resp, err := http.Get(apiURL)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Erro ao buscar anime"})
		}
		defer resp.Body.Close()

		var jikanResp JikanResponse
		if err := json.NewDecoder(resp.Body).Decode(&jikanResp); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Erro ao decodificar resposta"})
		}

		if len(jikanResp.Data) == 0 {
			return c.Status(404).JSON(fiber.Map{"message": "Anime não encontrado"})
		}

		return c.JSON(jikanResp.Data[0])
	}
}

func GetTopAnimePopular() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiURL := fmt.Sprintf("https://api.jikan.moe/v4/top/anime?filter=bypopularity&limit=25")

		resp, err := http.Get(apiURL)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": 500, "message": "Erro ao procurar animes populares"})
		}
		defer resp.Body.Close()

		var jikanResp JikanResponse

		if err := json.NewDecoder(resp.Body).Decode(&jikanResp); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": 500, "message": "Erro ao transformar resposta da api"})
		}

		if len(jikanResp.Data) == 0 {
			return c.Status(404).JSON(fiber.Map{"status": 404, "message": "Animes não encontrados"})
		}

		return c.Status(200).JSON(fiber.Map{"status": 200, "animes": jikanResp.Data})
	}
}

func GetTopAnimeAiring() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiURL := fmt.Sprintf("https://api.jikan.moe/v4/top/anime?filter=airing&limit=25")

		resp, err := http.Get(apiURL)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": 500, "message": "Erro ao procurar animes no ar"})
		}

		defer resp.Body.Close()

		var jikanResp JikanResponse

		fmt.Print(resp)

		if err = json.NewDecoder(resp.Body).Decode(&jikanResp); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": 500, "message": "Erro ao transformar resposta da API"})
		}

		if len(jikanResp.Data) == 0 {
			return c.Status(404).JSON(fiber.Map{"status": 404, "message": "Animes não encontrados"});
		}

		return c.Status(200).JSON(fiber.Map{"status": 200, "animes": jikanResp.Data});
	}
}

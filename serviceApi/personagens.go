package serviceapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/arcesti/animeApp/model"
	"fmt"
	"net/http"
	"encoding/json"
)

type JikanPersonagens struct {
	Data []model.Personagens `json:"data"`
}

func GetPersonagens () fiber.Handler {
	return func (c *fiber.Ctx) error {
		id := c.Params("id");
		apiURL := fmt.Sprintf("https://api.jikan.moe/v4/anime/%s/characters", id);

		resp, err := http.Get(apiURL);

		if err != nil {
			return c.Status(500).JSON(fiber.Map{ "message": "Erro ao buscar personagens" });
		}
		defer resp.Body.Close();

		var jikanResp JikanPersonagens;

		if err = json.NewDecoder(resp.Body).Decode(&jikanResp); err != nil {
			return c.Status(404).JSON(fiber.Map{ "message": "Erro ao tratar dados retornados" });
		}

		if len(jikanResp.Data) == 0 {
			return c.Status(404).JSON(fiber.Map{ "message": "Não foi encontrado nenhum personagem" });
		}

		return c.Status(200).JSON(fiber.Map{ "message": "sucesso ao recuperar personagens", "personagens": jikanResp });
	}
}
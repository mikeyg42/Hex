package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mikeyg42/Hex_demo/database"
	"github.com/mikeyg42/Hex_demo/models"
)

func ListMoves(c *fiber.Ctx) error {
	moves := []models.Move{}
	database.DB.Db.Find(&moves)

	return c.Status(200).JSON(moves)
}

func MakeMove(c *fiber.Ctx) error {
	newMove := new(models.Move)
	if err := c.BodyParser(newMove); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&newMove)

	return c.Status(200).JSON(newMove)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iLahh/shortener-url-service/database"
)

type statsResponse struct {
	ShortCode  string `json:"short_code"`
	LongURL    string `json:"long_url"`
	Expiry     int64  `json:"expiry_hours"`
	ClickCount int64  `json:"click_count"`
	CreatedAt  string `json:"created_at"`
}

func GetStats(c *fiber.Ctx) error {
	shortCode := c.Params("url")

	var stats statsResponse
	err := database.DB.QueryRow(
		`SELECT short_code, long_url, expiry, click_count, created_at
         FROM urls WHERE short_code = $1`,
		shortCode,
	).Scan(&stats.ShortCode, &stats.LongURL, &stats.Expiry, &stats.ClickCount, &stats.CreatedAt)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short URL not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}

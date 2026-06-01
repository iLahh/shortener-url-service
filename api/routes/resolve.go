package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/iLahh/shortener-url-service/database"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		var longURL string
		dbErr := database.DB.QueryRow(
			"SELECT long_url FROM urls WHERE short_code = $1", url,
		).Scan(&longURL)

		if dbErr != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "short not found in the database",
			})
		}

		value = longURL

	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	database.DB.Exec(
		"UPDATE urls SET click_count = click_count + 1 WHERE short_code = $1", url,
	)

	return c.Redirect(value, 301)
}

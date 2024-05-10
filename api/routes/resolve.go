package routes

import (
	"github.com/Yash-sudo-web/urlshortnerredisgolang/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	rdb := database.CreateClient(0)
	defer rdb.Close()

	val, err := rdb.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "URL not found",
		})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{url: val})
}

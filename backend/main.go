package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  app.Get("/health", func(c *fiber.Ctx) error {
    return c.SendString("OK")
  })

  app.Post("/input", func(c *fiber.Ctx) error {
    type Body struct {
      Input string `json:"input"`
    }
    var body Body
    if err := c.BodyParser(&body); err != nil {
      return err
    }
    return c.JSON(body)
  })

  app.Listen(":3000")
}

package api

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func h(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberHandlerFunc(h)(w, r)
}

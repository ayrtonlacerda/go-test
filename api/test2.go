package api

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// GET /api/test2
func get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "get2"})
}

// POST /api/test2
func post(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "post2"})
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		adaptor.FiberHandlerFunc(get)(w, r)
	}
	if r.Method == "POST" {
		adaptor.FiberHandlerFunc(post)(w, r)
	}
}

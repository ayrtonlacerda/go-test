package api

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email string
	Pass  string
}

// GET /api/test
func getTest(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "get"})
}

// POST /api/test
func postTest(c *fiber.Ctx) error {
	user := new(User)
	err := c.BodyParser(user)

	if err != nil {
		return c.Status(405).JSON(fiber.Map{"msg": "não foi possivel add um novo post"})
	}

	return c.JSON(user)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		adaptor.FiberHandlerFunc(getTest)(w, r)
	}
	if r.Method == "POST" {
		adaptor.FiberHandlerFunc(postTest)(w, r)
	}
}

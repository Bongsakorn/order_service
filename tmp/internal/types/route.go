package types

import "github.com/gofiber/fiber/v2"

type Route struct {
	Version     string
	Path        string
	Middlewares []fiber.Handler
	Children    []Children
}

type Children struct {
	Middlewares []fiber.Handler
	Handler     fiber.Handler
	Method      string
	Path        string
}

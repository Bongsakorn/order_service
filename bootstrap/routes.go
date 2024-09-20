package bootstrap

import (
	"fmt"
	"order_service/internal/config"

	"order_service/internal/types"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func Setup(app *fiber.App, config *config.ConfigService, routes []types.Route) {
	println("Initializing app...")

	SetAppMiddleware(app)
	api := app.Group("/api")
	lo.ForEach(routes, func(route types.Route, index int) {
		version := api.Group("/v" + route.Version)

		module := version.Group(route.Path, route.Middlewares...)
		lo.ForEach(route.Children, func(child types.Children, index int) {
			module.Add(child.Method, child.Path, child.Handler)
			println(fmt.Sprintf("Initializing {%s %s%s} added complete", child.Method, route.Path, child.Path))
		})
	})

	err := app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		panic(err)
	}
	println("App is running on port " + config.Port)
}

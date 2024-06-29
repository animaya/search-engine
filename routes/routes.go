package routes

import (
	"animaya/search-engine/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHadler := templ.Handler(component)

	for _, o := range options {
		o(componentHadler)
	}

	return adaptor.HTTPHandler(componentHadler)(c)
}

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, views.Home())
	})
}

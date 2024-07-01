package routes

import (
	"animaya/search-engine/views"
	"fmt"

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

type settingsform struct {
	Amount   int  `form:"amount"`
	SearchOn bool `form:"seachOn"`
	AddNew   bool `form:"addNew"`
}

func SetRoutes(app *fiber.App) {
	app.Get("/", LoginHandler)
	app.Post("/", func(c *fiber.Ctx) error {
		input := settingsform{}

		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error: Something went wrong</h2>")
		}
		fmt.Println(input)
		return c.SendStatus(200)

	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return render(c, views.Login())
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		input := loginform{}

		if err := c.BodyParser(&input); err != nil {
			return c.SendString("<h2>Error: Something went wrong</h2>")
		}

		return c.SendStatus(200)

	})
}

package routes

import (
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHadler := templ.Handler(component)

	for _, o := range options {
		o(componentHadler)
	}

	return adaptor.HTTPHandler(componentHadler)(c)
}

func SetRoutes(app *fiber.App) {

	app.Get("/login", LoginHandler)
	app.Post("/login", LoginPostHandler)
	app.Post("/logout", LogoutHandler)
	// app.Get("/create", func(c *fiber.Ctx) error {
	// 	u := &db.User{}
	// 	u.CreateAdmin()
	// 	return c.SendString("created")
	// })
	app.Post("/search", HandleSearch)
	app.Use("/search", cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	app.Get("/delete-admin", DeleteAdminUserHandler)

	app.Get("/", AuthMiddleware, DashboardHandler)
	app.Post("/", AuthMiddleware, DashboardPostHandler)
}

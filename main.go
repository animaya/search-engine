package main

import (
	"animaya/search-engine/db"
	"animaya/search-engine/routes"
	"animaya/search-engine/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))
	fmt.Println("SECRET_KEY:", os.Getenv("SECRET_KEY"))
	env := godotenv.Load()

	if env != nil {
		panic("cannot find environment variables")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(compress.New())
	db.InitDB()
	routes.SetRoutes(app)
	utils.StartCronJobs()

	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	app.Shutdown()

	fmt.Println("shutting down server")
}

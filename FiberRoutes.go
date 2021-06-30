package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

var app = fiber.New()

func FiberRoutes() {
	// Default config
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
	app.Get("/id", FiberIdGET)
	app.Get("/config/:id", FiberConfigGET)
	app.Post("/config/:id", FiberConfigPOST)
	app.Get("/result/:id", FiberResultGET)
	app.Get("/log/:id", FiberLogGET)
	port, ok := os.LookupEnv("PORT") // For Heroku or Elastic
	if !ok {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

var app = fiber.New()

func FiberRoutes() {
	app.Get("/id", FiberIdGET)
	app.Get("/config/:id", FiberConfigGET)
	app.Post("/config/:id", FiberConfigPOST)
	app.Get("/result/:id", FiberResultGET)
	log.Fatal(app.Listen(":8080"))
}

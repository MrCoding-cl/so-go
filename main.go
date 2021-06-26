package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

var server = createServer()

func main() {
	app := fiber.New()
	app.Use("/ws", FibberMiddleware)
	app.Get("/ws/:id", websocket.New(Socket))
	app.Get("/id", FiberIdGET)
	app.Get("/config/:id", FiberConfigGET)
	app.Post("/config/:id", FiberConfigPOST)
	app.Get("/result/:id", FiberResultGET)

	log.Fatal(app.Listen(":8080"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}

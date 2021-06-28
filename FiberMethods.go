package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"net/http"
	"strconv"
)

func FiberMiddleware(c *fiber.Ctx) error {
	// IsWebSocketUpgrade returns true if the passenger
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
func FiberIdGET(c *fiber.Ctx) error {
	return c.SendString(strconv.Itoa(server.add_client(&server)))
}
func FiberConfigGET(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	client, ok := server.clients[id]
	if ok {
		return c.JSON(client.Config)
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}
func FiberConfigPOST(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	client, ok := server.clients[id]
	if ok {
		err := c.BodyParser(&client.Config)
		if err != nil {
			return c.SendStatus(http.StatusNotAcceptable)
		}
		return c.SendStatus(http.StatusOK)
	}
	return c.SendStatus(http.StatusNotFound)
}
func FiberResultGET(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	client, ok := server.clients[id]
	if ok {
		err = getRoutine(client)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		if client.World == nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(&client.World)
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}

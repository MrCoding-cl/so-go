package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
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
func FiberSocket(c *websocket.Conn) {
	// c.Locals is added to the *websocket.Conn
	//log.Println(c.Locals("allowed"))  // true
	//log.Println(c.Params("id"))       // 123
	//log.Println(c.Query("v"))         // 1.0
	//log.Println(c.Cookies("session")) // ""
	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		_ = c.WriteMessage(http.StatusUnprocessableEntity, []byte("Bad route"))
		return
	}
	client, ok := server.clients[id]
	if !ok {
		_ = c.WriteMessage(http.StatusNotFound, []byte("Client ID not found"))
		return
	}
	var (
		msg []byte
	)
	for {
		if _, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		str := string(msg)
		if str == "start" {
			client.World = createWorld(12000)
			client.World.socket = c
			switch client.Config.RunType {
			case 0:
				err := morningRoutine(client.World)
				if err != nil {
					_ = c.WriteMessage(http.StatusInternalServerError, []byte(err.Error()))
					return
				}
			case 1:
				err := afternoonRoutine(client.World)
				if err != nil {
					_ = c.WriteMessage(http.StatusInternalServerError, []byte(err.Error()))
					return
				}
			case 2:
				err := nightRoutine(client.World)
				if err != nil {
					_ = c.WriteMessage(http.StatusInternalServerError, []byte(err.Error()))
					return
				}
			case 3:
				randomRoutine(client.World)
			case 4:
				err := CustomRoutine(client.World, client)
				if err != nil {
					_ = c.WriteMessage(http.StatusInternalServerError, []byte(err.Error()))
					return
				}
			default:
				_ = c.WriteMessage(http.StatusNotAcceptable, []byte("Not Implemented Yet or something is wrong"))
				return
			}
			if client.Config.Pram {
				client.World.runWithPram(client.World)
			} else {
				client.World.runwWithoutPram(client.World)
			}
		}
	}

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
		if client.World == nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(&client.World)
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}

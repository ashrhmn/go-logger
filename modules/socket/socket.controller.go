package socket

import (
	"github.com/ashrhmn/go-logger/guards"
	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocketController struct {
	connectionPool storage.WsPool
}

func newSocketController(connectionPool storage.WsPool) SocketController {
	return SocketController{
		connectionPool: connectionPool,
	}
}

func (sc SocketController) RegisterRoutes(app fiber.Router) {
	router := app.Group("/socket")
	router.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Get(
		"/ws/logs",
		guards.AnyLoggedIn(""),
		websocket.New(func(c *websocket.Conn) {
			defer c.Close()
			for {
				id := primitive.NewObjectID().Hex()
				sc.connectionPool.Connections[id] = c
				if _, _, err := c.ReadMessage(); err != nil {
					delete(sc.connectionPool.Connections, id)
					log.Error("read:", err)
					break
				}
			}
		}),
	)
}

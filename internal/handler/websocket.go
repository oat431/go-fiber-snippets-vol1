package handler

import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3/log"
)

// WebSocketChat is a minimal echo-style WebSocket handler.
// Replace the body with your own protocol logic.
func WebSocketChat(c *websocket.Conn) {
	log.Infof("ws connected (host=%s)", c.Locals("Host"))
	defer log.Infof("ws disconnected (host=%s)", c.Locals("Host"))

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				log.Warnf("ws read: %v", err)
			}
			break
		}
		log.Infof("ws recv: %s", msg)

		reply := []byte("echo: " + string(msg))
		if err := c.WriteMessage(mt, reply); err != nil {
			log.Warnf("ws write: %v", err)
			break
		}
	}
}

package controller

import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3/log"
)

func WebSocketExample(c *websocket.Conn) {
	log.Info("WebSocket connection established ", c.Locals("Host"))
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Info("read:", err)
			break
		}
		log.Info("recv: %s", msg)
		reply := "from computer: " + string(msg)
		err = c.WriteMessage(mt, []byte(reply))
		if err != nil {
			log.Info("write:", err)
			break
		}
	}
}

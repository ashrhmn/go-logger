package storage

import "github.com/gofiber/contrib/websocket"

type WsPool struct {
	Connections map[string]*websocket.Conn
}

func newWsPool() WsPool {
	return WsPool{
		Connections: make(map[string]*websocket.Conn),
	}
}

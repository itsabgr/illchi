package broker

import "github.com/itsabgr/fastintmap"
import "github.com/fasthttp/websocket"

type webSocketMap struct {
	intMap fastintmap.Map
}

func (m *webSocketMap) Get(id uintptr) *websocket.Conn {
	conn, found := m.intMap.Get(id)
	if !found {
		return nil
	}
	return conn.(*websocket.Conn)
}

func (m *webSocketMap) Add(id uintptr, conn *websocket.Conn) bool {
	return m.intMap.Add(id, conn)
}

func (m *webSocketMap) Delete(id uintptr) {
	m.intMap.Delete(id)
}

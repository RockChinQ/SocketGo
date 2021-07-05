package conn

import (
	"net"
	"sync"
)

var SocketPool = make(map[int]*Handler)

var sid_index = 0
var Lock sync.Mutex

func AddHandler(conn net.Conn, as string) Handler {
	//Lock sid_index&pool,set sid of the new handler,sid_index+=1
	//Golang is fucking elegant.
	Lock.Lock()
	handler := NewHandler(conn, as)
	SocketPool[sid_index] = &handler
	sid_index++
	Lock.Unlock()
	return handler
}

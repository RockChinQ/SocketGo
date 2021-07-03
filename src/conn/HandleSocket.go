package conn

import (
	"net"
	"time"
)

type Handler struct {
	SID    int    //unique id
	As     string //client or server
	Socket net.Conn
	ConnT  time.Time //timestamp of connection establishing
	UpV    int       //upload speed
	DownV  int       //download speed
	UpD    uint32    //amount of upload data
	DownD  uint32    //amount of download data
	Recv   string    //received data
}

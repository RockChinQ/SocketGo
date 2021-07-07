package server

import (
	"SocketGo/src/conn"
	"SocketGo/src/util"
	"net"
	"strconv"
	"time"
)

type Listener struct {
	LID    int
	Lsn    net.Listener
	OpenT  time.Time //the time when the port open
	Status string
}

func NewListener(lsn net.Listener) Listener {
	var l Listener
	l.LID = lid_index
	l.Lsn = lsn
	l.OpenT = time.Now()
	l.Status = "opened"
	return l
}

func (l *Listener) AcceptConn() {
	for {
		c, err := l.Lsn.Accept()
		if err != nil {
			l.Status = "stopped"
			util.DebugMsg("Lsn-"+strconv.Itoa(l.LID), "err:"+err.Error())
			break
		}
		conn.AddHandler(c, "server")
	}
}

func (l *Listener) Dispose() error {
	return l.Lsn.Close()
}

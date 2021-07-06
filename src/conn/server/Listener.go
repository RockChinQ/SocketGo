package server

import (
	"net"
	"time"
)

type Listener struct {
	LID   int
	Lsn   net.Listener
	OpenT time.Time //the time when the port open
}

func NewListener(lsn net.Listener) Listener {
	var l Listener
	l.LID = lid_index
	l.Lsn = lsn
	l.OpenT = time.Now()
	return l
}

func (l *Listener) Dispose() error {
	return l.Lsn.Close()
}

package client

import (
	"net"
	"time"
)

var Timeout = 10000

func Make(addr string) (net.Conn, error) {
	conn1, err := net.DialTimeout("tcp", addr, time.Millisecond*time.Duration(Timeout))
	if err != nil {
		return nil, err
	}
	return conn1, nil
}

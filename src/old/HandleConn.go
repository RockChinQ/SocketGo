package main

import (
	"bufio"
	"fmt"
	"net"
)

var conn net.Conn

func MakeConn(address string) (err error) {
	conn1, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	conn = conn1
	return nil
}
func ReadMsg() {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("[ERR]%v\n", err)
			conn.Close()
			conn = nil
			fmt.Printf("[CONN]Close conn actively.\n")
			return
		}
		fmt.Printf("[RECV]%v\n", string(buf[:n]))
	}
}
func WriteMsg(msg string) (err error) {
	writer := bufio.NewWriter(conn)
	_, err = writer.Write([]byte(msg + "\n"))
	writer.Flush()
	fmt.Printf("[SEND]%v\n", msg)
	return err
}

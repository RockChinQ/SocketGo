package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

/**
键盘输入读取
*/
var stdin *bufio.Reader

func main() {
	stdin = bufio.NewReader(os.Stdin)
	go ReadStd()

	wg.Add(1)
	wg.Wait()
}
func ReadStd() {
	for {
		word, _, _ := stdin.ReadLine()
		spt := strings.Split(string(word), " ")
		switch spt[0] {
		case "/conn": // /conn addr:port
			err := MakeConn(spt[1])
			if err != nil {
				fmt.Printf("[ERR]%v\n", err)
				continue
			}
			fmt.Printf("[CONN]Conn established.\n")
			go ReadMsg()
			continue
		default:
			if conn == nil {
				fmt.Printf("[CONN]No conn.\n")
				continue
			}
			err := WriteMsg(string(word))
			if err != nil {
				fmt.Printf("[ERR]%v\n", err)
			}
			continue
		}
	}
}

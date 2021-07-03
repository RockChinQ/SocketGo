package main

import (
	"fmt"
	"strconv"

	"github.com/nsf/termbox-go"
)

func TermboxTest() {
	termbox.Init()
	termbox.SetCursor(0, 0)
	for {
		eve := termbox.PollEvent()
		if eve.Err != nil {
			panic(eve.Err)
		}
		fmt.Println(strconv.Itoa(int(eve.Key)) + " ch=" + string(eve.Ch))
	}
}

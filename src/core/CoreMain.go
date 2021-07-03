package main

import (
	"SocketGo/src/stdin"
	"SocketGo/src/util"
	"sync"
)

/*
Command
!server
	!server                                       list all port listeners
	!server [open <port:int>|close <port:int>]    open/close port listener
!client
	!client conn <addr:string>:<port:int>         make a new socket connected to specific port
!list
	!list                                         list all sockets to ports and from remote hosts
	!list [client|server]                         list all socket to port or from remote hosts
!kill
	!kill <SID:int>                               kill specific socket
!focus
	!focus <SID:int>                              focus specific socket
*/

var wg sync.WaitGroup

//Init program
func main() {
	util.InitTermbox()

	util.DebugMsg("Main", "Booting.")
	stdin.RegisterFuns()

	wg.Add(1)

	go stdin.Loop()

	wg.Wait()
}

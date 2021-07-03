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
!exit
	!exit                                         dispose resources and exit
*/

var wg sync.WaitGroup

//Init program
func main() {
	Welcome()
	util.DebugMsg("Main", "Booting.")
	stdin.RegisterFuns()

	stdin.PutPrompt()
	wg.Add(1)

	go stdin.Loop()

	wg.Wait()
}

func Welcome() {
	util.Sayln("      _________              __           __     ________")
	util.Sayln("     /   _____/ ____   ____ |  | __ _____/  |_  /  _____/  ____")
	util.Sayln("     \\_____  \\ /  _ \\_/ ___\\|  |/ // __ \\   __\\/   \\  ___ /  _ \\")
	util.Sayln("     /        (  <_> )  \\___|    <\\  ___/|  |  \\    \\_\\  (  <_> )")
	util.Sayln("    /_______  /\\____/ \\___  >__|_  \\___  >__|   \\______  /\\____/")
	util.Sayln("            \\/            \\/     \\/    \\/              \\/")
	util.Sayln("\nMade by Rock Chin,just for fun. See:https://github.com/RockChinQ/SocketGo\nType \"!help\" to get started. Command history can only work on Windows.\n")
}

package fun

import (
	"SocketGo/src/util"
	"os"
)

//Dispose resources and exit:all sockets,opend ports
func FuncExit(args []string, cmd string) map[string]string {
	DisposeAll()
	util.SaySub("FuncExit", "Exiting.")
	os.Exit(0)
	return EmptyMap()
}

func DisposeAll() {
	//kill all conns
	FuncKill([]string{"!kill", "all"}, "!kill all")
	//close all ports
	FuncServer([]string{"!server", "close", "all"}, "!server close all")
}

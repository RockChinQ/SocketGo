package fun

import "SocketGo/src/util"

func FuncHelp(args []string, cmd string) map[string]string {
	util.SaySub("FuncHelp", "Name\tParam\tDescription")
	util.SaySub("FuncHelp", "!help\t\tDisplay this message")
	return EmptyMap()
}

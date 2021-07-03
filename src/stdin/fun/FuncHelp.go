package fun

import "SocketGo/src/util"

func FuncHelp(args []string, cmd string) map[string]string {
	util.SaySub("Help", "Name\tParam\tDescription")
	util.SaySub("Help", "!help\t\tDisplay this message")
	return make(map[string]string)
}

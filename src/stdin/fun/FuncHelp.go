package fun

import "SocketGo/src/util"

func FuncHelp(args []string, cmd string) map[string]string {
	util.SaySub("Help", "Name\tParam\tDescription")
	return make(map[string]string)
}

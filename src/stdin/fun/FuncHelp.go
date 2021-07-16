package fun

import (
	"SocketGo/src/model"
	"SocketGo/src/util"
)

func FuncHelp(info *model.ExecInfo) {
	util.SaySub("FuncHelp", "Name\tParam\tDescription")
	util.SaySub("FuncHelp", "!help\t\tDisplay this message")
}

package fun

import (
	"SocketGo/src/model"
)

func FuncHelp(info *model.ExecInfo) {
	info.SaySub("FuncHelp", "Name\tParam\tDescription")
	info.SaySub("FuncHelp", "!help\t\tDisplay this message")
}

package fun

import (
	"SocketGo/src/model"
)

func FuncHelp(info *model.ExecInfo) {
	info.SaySub("FuncHelp", "Name\tOperation&Params          \tDescription")
	info.SaySub("FuncHelp", "help\t                           \tdisplay this message")
	info.SaySub("FuncHelp", "echo\t<msg:string>               \techo message")
	info.SaySub("FuncHelp", "server\t                           \tlist all open ports")
	info.SaySub("FuncHelp", "~     \topen/close <port:int>       \topen or close specific port")
	info.SaySub("FuncHelp", "client\tconn <addr:string>:<port:int>\tconnect to specific port on specific host")
	info.SaySub("FuncHelp", "~     \ttimeout <mills:int>       \tset timeout of establishing a conn")
	info.SaySub("FuncHelp", "list\t                          \tlist all established conns")
	info.SaySub("FuncHelp", "~   \tclient/server             \tlist all conn of client or server")
	info.SaySub("FuncHelp", "kill\t<SID:int>                 \tkill specific conn whose SID is the provided one")
	info.SaySub("FuncHelp", "~   \tall                       \tkill all conns")
	info.SaySub("FuncHelp", "exit\t                           \tclose all ports and kill all conns then exit program")
}

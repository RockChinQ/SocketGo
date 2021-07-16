package fun

import (
	"SocketGo/src/conn/server"
	"SocketGo/src/model"
	"SocketGo/src/util"
	"strconv"
)

func FuncServer(info *model.ExecInfo) {
	if len(info.Args) == 1 { //no extra params,list ports
		info.SaySub("FuncServer", "LID\tAddr:Port\tTime")
		for k, v := range server.ListenerList {
			info.SaySub("FuncServer", strconv.Itoa(k)+"\t"+v.Lsn.Addr().String()+"/"+v.Lsn.Addr().Network()+"\t"+util.GetTimeStr(v.OpenT))
		}
		info.SaySub("FuncServer", "Done,count:"+strconv.Itoa(len(server.ListenerList)))
	} else { //with params
		switch info.Args[1] {
		case "open": //open specific port
			if len(info.Args) < 3 {
				info.SaySub("FuncServer", "err:Please provide port to open.")
				info.Error("err:Please provide port to open.")
				return
			}
			port, err := strconv.Atoi(info.Args[2])
			if err != nil {
				info.SaySub("FuncServer", "err:args[2] is not a number.")
				info.Error("err:args[2] is not a number.")
				return
			}
			lsn, err := server.OpenListener(port)
			if err != nil {
				info.SaySub("FuncServer", "err:Cannot open port("+strconv.Itoa(port)+"):"+err.Error())
				info.Error("err:Cannot open port(" + strconv.Itoa(port) + "):" + err.Error())
				return
			}
			l := server.AddListener(lsn)
			info.SaySub("FuncServer", "Successfully opened port:"+l.Lsn.Addr().String()+"/"+l.Lsn.Addr().Network()+" LID="+strconv.Itoa(l.LID))
		case "close": //close specific port
			if len(info.Args) < 3 {
				info.SaySub("FuncServer", "err:Please provide port to close.")
				info.Error("err:Please provide port to close.")
				return
			}
			if info.Args[2] == "all" {
				//lock the lpool
				server.Lock.Lock()
				//loop all element,call dispose func of each lsn
				removed := make(map[int]*server.Listener)
				for k, v := range server.ListenerList {
					err := v.Dispose()
					if err != nil {
						info.SaySub("FuncServer", "Unable to close port:"+v.Lsn.Addr().String()+" :"+err.Error())
						continue
					} else {
						removed[k] = v //add to temp map
					}
				}
				//remove element from lsnpool
				for k, v := range removed {
					info.SaySub("FuncServer", "Successfully closed:SID="+itoa(k)+" port="+v.Lsn.Addr().String())
					delete(server.ListenerList, k)
				}
				server.Lock.Unlock()
			} else {
				port, err := strconv.Atoi(info.Args[2])
				if err != nil {
					info.SaySub("FuncServer", "err:args[2] is not a number.")
					info.Error("err:args[2] is not a number.")
					return
				}
				s := server.DisposeListener(port)
				if s {
					info.SaySub("FuncServer", "Successfully closed port:"+strconv.Itoa(port))
				} else {
					info.SaySub("FuncServer", "Unable to close port:"+strconv.Itoa(port))
				}
			}
		default:
			info.SaySub("FuncServer", "err:No such operation:"+info.Args[1])
			info.Error("err:No such operation:" + info.Args[1])
			return
		}
	}
}

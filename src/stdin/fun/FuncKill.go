package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/model"
	"strconv"
)

func FuncKill(info *model.ExecInfo) {

	if len(info.Args) < 2 {
		info.SaySub("FuncKill", "err:Please provide SID/\"all\" to kill specific/all conn.")
		info.Error("err:Please provide SID/\"all\" to kill specific/all conn.")
		return
	}
	if info.Args[1] == "all" {
		conn.Lock.Lock()
		removed := make(map[int]*conn.Handler)
		for k, v := range conn.SocketPool {
			err := v.Dispose()
			if err != nil {
				info.SaySub("FuncKill", "err:Occurs while killing conn SID="+strconv.Itoa(k))
				continue
			}
			removed[k] = v
		}
		for k := range removed {
			delete(conn.SocketPool, k)
			info.SaySub("FuncKill", "Successfully killed conn SID="+strconv.Itoa(k))
		}
		conn.Lock.Unlock()
	} else { //kill specific
		v, err := strconv.Atoi(info.Args[1])
		if err != nil {
			info.SaySub("FuncKill", "err:args[1] is not \"all\" or a number.")
			info.Error("err:args[1] is not \"all\" or a number.")
			return
		}
		conn.Lock.Lock()
		err = conn.SocketPool[v].Dispose()
		if err != nil {
			info.SaySub("FuncKill", "err:"+err.Error())
			conn.Lock.Unlock()
			info.Error("err:" + err.Error())
			return
		} else {
			delete(conn.SocketPool, v)
			conn.Lock.Unlock()
			info.SaySub("FuncKill", "Successfully kill conn:SID="+strconv.Itoa(v))
		}
	}
}

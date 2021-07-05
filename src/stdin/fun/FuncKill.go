package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/util"
	"strconv"
)

func FuncKill(args []string, cmd string) map[string]string {

	if len(args) < 2 {
		util.SaySub("Kill", "err:Please provide SID/\"all\" to kill specific/all conn.")
		return EmptyMap()
	}
	if args[1] == "all" {
		conn.Lock.Lock()
		removed := make(map[int]*conn.Handler)
		for k, v := range conn.SocketPool {
			err := v.Dispose()
			if err != nil {
				util.SaySub("Kill", "err:Occurs while killing conn SID="+strconv.Itoa(k))
				continue
			}
			removed[k] = v
		}
		for k, _ := range removed {
			delete(conn.SocketPool, k)
			util.SaySub("Kill", "Successfully killed conn SID="+strconv.Itoa(k))
		}
		conn.Lock.Unlock()
	} else { //kill specific
		v, err := strconv.Atoi(args[1])
		if err != nil {
			util.SaySub("Kill", "err:args[1] is not \"all\" or a number.")
			return EmptyMap()
		}
		conn.Lock.Lock()
		conn.SocketPool[v].Dispose()
		delete(conn.SocketPool, v)
		conn.Lock.Unlock()
	}
	return EmptyMap()
}

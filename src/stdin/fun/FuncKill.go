package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/util"
	"strconv"
)

func FuncKill(args []string, cmd string) map[string]string {

	if len(args) < 2 {
		util.SaySub("FuncKill", "err:Please provide SID/\"all\" to kill specific/all conn.")
		return ErrMap("err:Please provide SID/\"all\" to kill specific/all conn.")
	}
	if args[1] == "all" {
		conn.Lock.Lock()
		removed := make(map[int]*conn.Handler)
		for k, v := range conn.SocketPool {
			err := v.Dispose()
			if err != nil {
				util.SaySub("FuncKill", "err:Occurs while killing conn SID="+strconv.Itoa(k))
				continue
			}
			removed[k] = v
		}
		for k := range removed {
			delete(conn.SocketPool, k)
			util.SaySub("FuncKill", "Successfully killed conn SID="+strconv.Itoa(k))
		}
		conn.Lock.Unlock()
	} else { //kill specific
		v, err := strconv.Atoi(args[1])
		if err != nil {
			util.SaySub("FuncKill", "err:args[1] is not \"all\" or a number.")
			return ErrMap("err:args[1] is not \"all\" or a number.")
		}
		conn.Lock.Lock()
		err = conn.SocketPool[v].Dispose()
		if err != nil {
			util.SaySub("FuncKill", "err:"+err.Error())
			conn.Lock.Unlock()
			return ErrMap("err:" + err.Error())
		} else {
			delete(conn.SocketPool, v)
			conn.Lock.Unlock()
			util.SaySub("FuncKill", "Successfully kill conn:SID="+strconv.Itoa(v))
		}
	}
	return NoErrMap()
}

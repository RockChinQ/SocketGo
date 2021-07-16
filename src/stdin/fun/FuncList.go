package fun

import (
	"SocketGo/src/conn"
	"SocketGo/src/model"
	"SocketGo/src/util"
	"strconv"
)

func FuncList(info *model.ExecInfo) {
	if len(info.Args) == 1 { //all connection
		conn.Lock.Lock()
		util.SaySub("FuncList", "SID\tcreator\tconnTime\tspeed(up/down)\tdata(up/down)\tstatus")
		for k, v := range conn.SocketPool {
			util.SaySub("FuncList", itoa(k)+"\t"+v.As+"\t"+util.GetTimeStr(v.ConnT)+"\t"+itoa(v.UpV)+"/"+itoa(v.DownV)+"\t"+itoa32(v.UpD)+"/"+itoa32(v.DownD)+"\t"+v.Status)
		}
		util.SaySub("FuncList", "Done,count:"+strconv.Itoa(len(conn.SocketPool)))
		conn.Lock.Unlock()
	} else {
		if info.Args[1] == "client" || info.Args[1] == "server" {
			conn.Lock.Lock()
			count := 0
			util.SaySub("FuncList", "SID\tcreator\tconnTime\tspeed(up/down)\tdata(up/down)\tstatus")
			for k, v := range conn.SocketPool {
				if v.As == info.Args[1] {
					count++
					util.SaySub("FuncList", itoa(k)+"\t"+v.As+"\t"+util.GetTimeStr(v.ConnT)+"\t"+itoa(v.UpV)+"/"+itoa(v.DownV)+"\t"+itoa32(v.UpD)+"/"+itoa32(v.DownD)+"\t"+v.Status)
				}
			}
			util.SaySub("FuncList", "Done,count:"+strconv.Itoa(count))
			conn.Lock.Unlock()
		} else {
			util.SaySub("FuncList", "err:args[1] should be \"client\" or \"server\".")
			info.Error("err:args[1] should be \"client\" or \"server\".")
		}
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
func itoa32(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func NoErrMap() map[string]string {
	m := make(map[string]string)
	m["error"] = "NULL"
	return m
}
func ErrMap(err string) map[string]string {
	m := make(map[string]string)
	m["error"] = err
	return m
}

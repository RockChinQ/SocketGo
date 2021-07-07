package server

import (
	"net"
	"strconv"
	"sync"
)

var ListenerList = make(map[int]*Listener)
var lid_index = 0
var Lock sync.Mutex

//add listener to map
func AddListener(lsn net.Listener) Listener {
	Lock.Lock()
	listener := NewListener(lsn)
	ListenerList[lid_index] = &listener
	lid_index++
	Lock.Unlock()
	go listener.AcceptConn()
	return listener
}

//open a listener
func OpenListener(port int) (net.Listener, error) {
	l, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	return l, nil
}

func DisposeListener(port int) bool {
	Lock.Lock()
	var lsn *Listener
	for _, v := range ListenerList {
		if v.Lsn.Addr().String() == "127.0.0.1:"+strconv.Itoa(port) {
			lsn = v
			break
		}
	}
	if lsn == nil {
		Lock.Unlock()
		return false
	} else {
		lsn.Dispose()
		delete(ListenerList, lsn.LID)
		Lock.Unlock()
		return true
	}
}

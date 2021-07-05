package conn

import (
	"net"
	"time"
)

type Handler struct {
	SID   int    //unique id
	As    string //client or server
	Conn  net.Conn
	ConnT time.Time //timestamp of connection establishing
	UpV   int       //upload speed
	DownV int       //download speed
	UpD   uint32    //amount of upload data
	DownD uint32    //amount of download data
	Recv  string    //received data
}

//Init new Handler with created Conn and As label.This method is not thread-safe.
func NewHandler(conn net.Conn, as string) Handler {
	var h Handler
	h.SID = sid_index
	h.As = as
	h.Conn = conn
	h.ConnT = time.Now()
	h.UpV = 0
	h.DownV = 0
	h.UpD = 0
	h.DownD = 0
	h.Recv = ""
	return h
}

func (h *Handler) Dispose() error {
	return h.Conn.Close()
}

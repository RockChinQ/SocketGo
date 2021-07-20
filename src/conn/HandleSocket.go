package conn

import (
	"net"
	"time"
)

type Handler struct {
	SID    int    //unique id
	As     string //client or server
	Conn   net.Conn
	ConnT  time.Time //timestamp of connection establishing
	UpV    int       //upload speed
	DownV  int       //download speed
	UpD    uint32    //amount of upload data
	DownD  uint32    //amount of download data
	Recv   *string   //received data
	Buf    *string   //store buffer msg
	Status string
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
	es := ""
	h.Recv = &es
	h.Buf = &es
	h.Status = "established"
	go h.Read()
	return h
}

func (h *Handler) Read() {
	for {
		s, err := ReadString(h)
		if err != nil {
			return
		}
		*h.Buf = *h.Buf + s
		*h.Recv = *h.Recv + s
	}
}

func (h *Handler) Dispose() error {
	return h.Conn.Close()
}

func ReadString(h *Handler) (string, error) {
	buf := make([]byte, 1024)
	n, err := h.Conn.Read(buf)
	if err != nil {
		return "", err
	}
	h.DownD += uint32(n)
	return string(buf[:n]), nil
}

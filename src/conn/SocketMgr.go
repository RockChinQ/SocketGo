package conn

var SocketPool map[int]Handler

func AddHandler(h Handler) {

}

func RemoveHandler(sid int) {
	delete(SocketPool, sid)
}

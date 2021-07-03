package fun

import "os"

//Dispose resources and exit:all sockets,opend ports
func FuncExit(args []string, cmd string) map[string]string {
	DisposeAll()
	os.Exit(0)
	return make(map[string]string)
}

func DisposeAll() {

}

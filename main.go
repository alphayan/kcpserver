package main

import (
	"fmt"

	"github.com/xtaci/kcp-go"
	"net"
)

func main() {
	lis, err := kcp.ListenWithOptions(LISTENING_PORT, nil, 0, 0)
	if err != nil {
		fmt.Println("server line 41: ", err)
		return
	}

	for {
		conn, err := lis.AcceptKCP()
		if err != nil {
			fmt.Println("err accepting: " + err.Error())
			return
		}
		conn.SetNoDelay(1, 10, 2, 1)
		conn.SetWindowSize(1024, 1024)
		conn.SetACKNoDelay(true)
		go doServerStaff(conn)
	}
}

func doServerStaff(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 4096)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err reading")
			return
		}
		fmt.Println("received data: ", string(buf[:n]))
		conn.Write(buf[:n])
	}
}


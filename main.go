package main

import (
	"fmt"

	"github.com/xtaci/kcp-go"
)

func main() {
	lis, err := kcp.ListenWithOptions(LISTENING_PORT, nil, DATA_SHARD, PARITY_SHARD)
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
		go doServerStaff(conn)
	}
}

func doServerStaff(conn *kcp.UDPSession) {
	defer conn.Close()
	for {
		buf := make([]byte, 4096)
		conn.SetNoDelay(1, 10, 2, 1)
		conn.SetWindowSize(32, 32)
		conn.SetACKNoDelay(true)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err reading")
			return
		}
		fmt.Println("received data: ", string(buf[:n]))
		conn.Write(buf[:n])
	}
}


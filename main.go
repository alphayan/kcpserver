package main

import (
	"fmt"
	"net"

	"strings"

	"github.com/alphayan/kcp-go"
)

var addrs []string

func main() {
	lis, err := kcp.ListenWithOptions(LISTENING_PORT, nil, 0, 0)
	if err != nil {
		fmt.Println("server line 41: ", err)
		return
	}

	for {
		conn, err := lis.AcceptKCP()
		remote := conn.RemoteAddr().String()
		if len(addrs) > 0 {
			if !inSlice(remote, addrs) {
				addrs = append(addrs, remote)
			}
		} else {
			addrs = append(addrs, remote)
		}
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

func doServerStaff(conn *kcp.UDPSession) {
	defer conn.Close()
	for {
		buf := make([]byte, 4096)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err reading")
			return
		}
		fmt.Println("received data: ", string(buf[:n]))
		for _, v := range addrs {
			udpaddr, err := net.ResolveUDPAddr("udp", v)
			if err != nil {
				fmt.Println(err)
				return
			}
			conn.WriteTo(buf[:n], udpaddr)
		}

	}
}
func inSlice(s string, sl []string) bool {
	for _, v := range sl {
		if strings.EqualFold(s, v) {
			return true
		}
	}
	return false
}

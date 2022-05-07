package util

import (
	"fmt"
	"io"
	"net"
)

func Req(host string, port int, ch chan string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	defer conn.Close()

	b, err := io.ReadAll(conn)
	if err != nil {
		ch <- fmt.Sprintln(err)
	}
	ch <- string(b)
}

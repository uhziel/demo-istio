package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/uhziel/demo-istio/pkg/util"
)

var ratingsHost string
var ratingsPort int

func init() {
	flag.StringVar(&ratingsHost, "ratings-host", "localhost", "ratings host")
	flag.IntVar(&ratingsPort, "ratings-port", 2003, "ratings port")
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", ":2001")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			ch := make(chan string)
			go util.Req(ratingsHost, ratingsPort, ch)
			ratings := <-ch
			response := fmt.Sprintf("reviews v1 -> %s", ratings)
			c.Write([]byte(response))
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

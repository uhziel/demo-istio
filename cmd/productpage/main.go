package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/uhziel/demo-istio/pkg/util"
)

var reviewsHost string
var reviewsPort int

var detailsHost string
var detailsPort int

func init() {
	flag.StringVar(&reviewsHost, "reviews-host", "localhost", "reviews host")
	flag.IntVar(&reviewsPort, "reviews-port", 2001, "reviews port")

	flag.StringVar(&detailsHost, "details-host", "localhost", "details host")
	flag.IntVar(&detailsPort, "details-port", 2002, "details port")
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", ":2000")
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
			ch1 := make(chan string)
			go util.Req(reviewsHost, reviewsPort, ch1)
			ch2 := make(chan string)
			go util.Req(detailsHost, detailsPort, ch2)
			reviewsResp := <-ch1
			detailsResp := <-ch2
			respone := fmt.Sprintf("productpage -> %s, %s", detailsResp, reviewsResp)
			c.Write([]byte(respone))
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

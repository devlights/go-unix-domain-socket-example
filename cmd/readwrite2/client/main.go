package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/devlights/go-unix-domain-socket-example/pkg/message"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	values := []string{
		"hello world",
		"golang",
		"goroutine",
		"this program runs on crostini",
	}

	conn, err := net.Dial(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, v := range values {
		time.Sleep(1 * time.Second)

		m := &message.Echo{
			Length: len(v),
			Data:   []byte(v),
		}

		err = m.Write(conn)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[WRITE] ", m)

		err = m.Read(conn)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[READ ] ", m)
	}
}

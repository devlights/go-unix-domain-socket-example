package main

import (
	"fmt"
	"go-unix-domain-socket-example/pkg/message"
	"log"
	"net"
	"time"
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

	for _, v := range values {
		time.Sleep(1 * time.Second)

		conn, err := net.Dial(protocol, sockAddr)
		if err != nil {
			log.Fatal(err)
		}

		func() {
			defer conn.Close()

			m := &message.Echo{
				Length: len(v),
				Data:   []byte(v),
			}

			err = m.Write(conn)
			if err != nil {
				log.Fatal(err)
			}

			err = m.Read(conn)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", m)
		}()
	}
}

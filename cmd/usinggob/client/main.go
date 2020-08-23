package main

import (
	"encoding/gob"
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

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for _, v := range values {
		time.Sleep(1 * time.Second)

		m := &message.Echo{
			Length: len(v),
			Data:   []byte(v),
		}

		err = encoder.Encode(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[WRITE] ", m)

		err = decoder.Decode(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[READ ] ", m)
	}
}

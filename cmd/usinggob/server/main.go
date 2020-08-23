package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/devlights/go-unix-domain-socket-example/pkg/message"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	cleanup := func() {
		if _, err := os.Stat(sockAddr); err == nil {
			if err := os.RemoveAll(sockAddr); err != nil {
				log.Fatal(err)
			}
		}
	}

	cleanup()

	listener, err := net.Listen(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		fmt.Println("ctrl-c pressed!")
		close(quit)
		cleanup()
		os.Exit(0)
	}()

	fmt.Println("> Server launched")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(">>> accepted: ", conn.RemoteAddr().Network())
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)
	for {
		m := &message.Echo{}
		err := decoder.Decode(m)
		if err != nil {
			if err == io.EOF {
				fmt.Println("=== closed by client")
				break
			}

			log.Println(err)
			break
		}

		fmt.Println("[READ ] ", m)

		s := strings.ToUpper(string(m.Data))
		m.Length = len(s)
		m.Data = []byte(s)

		err = encoder.Encode(m)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("[WRITE] ", m)
	}
}

package main

import (
	"fmt"
	"go-unix-domain-socket-example/pkg/message"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	if _, err := os.Stat(sockAddr); err == nil {
		if err := os.RemoveAll(sockAddr); err != nil {
			log.Fatal(err)
		}
	}

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
		os.Exit(0)
	}()

	fmt.Println("> Server launched")
	for {
		fmt.Println("> wait...")

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

	m := &message.Echo{}
	err := m.Read(conn)
	if err != nil {
		log.Println(err)
		return
	}

	s := strings.ToUpper(string(m.Data))
	m.Length = len(s)
	m.Data = []byte(s)

	err = m.Write(conn)
	if err != nil {
		log.Println(err)
		return
	}
}

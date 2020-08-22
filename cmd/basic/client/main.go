package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)

		conn, err := net.Dial(protocol, sockAddr)
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte("hello world"))
		if err != nil {
			log.Fatal(err)
		}

		// サーバ側にクライアントからの書き込みが終わったことを
		// 無理やり伝えるためにWrite側のソケットを強制クローズ
		// (サンプル以外ではしてはいけない)
		err = conn.(*net.UnixConn).CloseWrite()
		if err != nil {
			log.Fatal(err)
		}

		b, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}
}

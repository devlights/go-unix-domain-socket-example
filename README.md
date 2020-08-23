# go-unix-domain-socket-example
Unix domain socket (UDS) example by golang

## How to Run

### basic

Open two terminals. the one is following:

```sh
$ cd cmd/basic/server
$ go run .
```

and another terminal is following:

```sh
$ cd cmd/basic/client
$ go run .
```

### read & write per connection

Open two terminals. the one is following:

```sh
$ cd cmd/readwrite/server
$ go run .
```

and another terminal is following:

```sh
$ cd cmd/readwrite/client
$ go run .
```

### N read & write on one connection

Open two terminals. the one is following:

```sh
$ cd cmd/readwrite2/server
$ go run .
```

and another terminal is following:

```sh
$ cd cmd/readwrite2/client
$ go run .
```

### Using encoding/gob package

Open two terminals. the one is following:

```sh
$ cd cmd/usinggob/server
$ go run .
```

and another terminal is following:

```sh
$ cd cmd/usinggob/client
$ go run .
```

## Monitor a local unix domain socket

use socat. install it if not exists.

```sh
$ sudo apt install -y socat
```

and launch server program then input following commands.

```sh
$ mv /tmp/echo.sock /tmp/echo.sock.original
$ socat -t100 -x -v UNIX-LISTEN:/tmp/echo.sock,mode=777,reuseaddr,fork UNIX-CONNECT:/tmp/echo.sock.original
```

### REFERENCES

[Can I monitor a local unix domain socket like tcpdump?](https://superuser.com/questions/484671/can-i-monitor-a-local-unix-domain-socket-like-tcpdump)


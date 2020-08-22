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


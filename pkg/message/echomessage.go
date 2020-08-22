package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type (
	Echo struct {
		Length int
		Data   []byte
	}
)

func (e *Echo) String() string {
	return fmt.Sprintf("Length[%02d] Data[%s]", e.Length, e.Data)
}

func (e *Echo) Write(c net.Conn) error {
	data := make([]byte, 0, 4+e.Length)

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(e.Length))
	data = append(data, buf...)

	w := bytes.Buffer{}
	err := binary.Write(&w, binary.BigEndian, e.Data)
	if err != nil {
		return err
	}

	data = append(data, w.Bytes()...)

	_, err = c.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (e *Echo) Read(c net.Conn) error {
	buf := make([]byte, 4)

	_, err := c.Read(buf)
	if err != nil {
		return err
	}

	byteCount := binary.BigEndian.Uint32(buf)
	e.Length = int(byteCount)
	e.Data = make([]byte, e.Length)

	_, err = c.Read(e.Data)
	if err != nil {
		return err
	}

	return nil
}

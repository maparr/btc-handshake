package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// readVerackMessage reads a verack message from the target node
func ReadVerackMessage(conn net.Conn) error {
	header := make([]byte, 24)
	if _, err := io.ReadFull(conn, header); err != nil {
		return err
	}

	magic := binary.LittleEndian.Uint32(header[0:4])
	if magic != binary.LittleEndian.Uint32([]byte{0xf9, 0xbe, 0xb4, 0xd9}) {
		return fmt.Errorf("invalid magic: %x", magic)
	}

	command := string(bytes.Trim(header[4:16], "\x00"))
	if command != CommandVerack {
		return fmt.Errorf("unexpected command: %s", command)
	}

	length := binary.LittleEndian.Uint32(header[16:20])
	if length != 0 {
		return fmt.Errorf("invalid verack length: %d", length)
	}

	fmt.Printf("Received message (%s)\n", command)
	return nil
}

package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// readMessage reads a message with the specified command from the target node
func ReadMessage(conn net.Conn, expectedCommand string) ([]byte, error) {
	header := make([]byte, 24)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}

	magic := binary.LittleEndian.Uint32(header[0:4])
	if magic != binary.LittleEndian.Uint32([]byte{0xf9, 0xbe, 0xb4, 0xd9}) {
		return nil, fmt.Errorf("invalid magic: %x", magic)
	}

	command := string(bytes.Trim(header[4:16], "\x00"))
	if command != expectedCommand {
		return nil, fmt.Errorf("unexpected command: %s", command)
	}

	length := binary.LittleEndian.Uint32(header[16:20])
	payload := make([]byte, length)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	fmt.Printf("Received message (%s): %x\n", command, payload)

	return payload, nil
}

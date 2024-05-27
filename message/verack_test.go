package message

import (
	"net"
	"testing"
)

func TestReadVerackMessage(t *testing.T) {
	// Simulate a TCP connection
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		// Send a valid verack message
		header := []byte{
			0xf9, 0xbe, 0xb4, 0xd9, // magic
			'v', 'e', 'r', 'a', 'c', 'k', 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // command
			0x00, 0x00, 0x00, 0x00, // length
			0x5d, 0xf6, 0xe0, 0xe2, // checksum
		}
		conn.Write(header)
	}()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	err = ReadVerackMessage(conn)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

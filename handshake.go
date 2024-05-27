package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/maparr/btc-handshake/message"
)

func main() {
	// Define command-line flags
	targetNode := flag.String("node", "127.0.0.1", "IP address of the target node")
	port := flag.String("port", message.DefaultPort, "Port of the target node")
	flag.Parse()

	fmt.Printf("Starting handshake with node at %s:%s\n", *targetNode, *port)

	conn, err := net.Dial("tcp", *targetNode+":"+*port)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
		fmt.Println("Connection closed.")
	}()

	fmt.Printf("Connected to node at %s:%s\n", *targetNode, *port)

	versionMsg := message.CreateVersionMessage(*targetNode)
	fmt.Println("Sending version message...")
	if err := sendMessage(conn, message.CommandVersion, versionMsg); err != nil {
		fmt.Printf("Failed to send version message: %v\n", err)
		return
	}
	fmt.Println("Version message sent.")

	fmt.Println("Waiting to receive version message...")
	versionPayload, err := message.ReadMessage(conn, message.CommandVersion)
	if err != nil {
		fmt.Printf("Failed to receive version message: %v\n", err)
		return
	}
	fmt.Println("Version message received.")

	receivedVersionMsg, err := message.DecodeVersionMessage(versionPayload)
	if err != nil {
		fmt.Printf("Failed to decode version message: %v\n", err)
		return
	}
	fmt.Printf("Received Version Message: %+v\n", receivedVersionMsg)

	fmt.Println("Waiting to receive verack message...")
	if err := message.ReadVerackMessage(conn); err != nil {
		fmt.Printf("Failed to receive verack: %v\n", err)
		return
	}
	fmt.Println("Verack message received.")

	fmt.Println("Handshake successful")
}

// sendMessage sends a message to the target node
func sendMessage(conn net.Conn, command string, payload interface{}) error {
	payloadBytes, err := message.EncodeVersionMessage(payload.(message.VersionMessage))
	if err != nil {
		return err
	}

	fmt.Printf("Payload to send (%s): %x\n", command, payloadBytes)

	header := message.CreateMessageHeader(command, payloadBytes)
	fmt.Printf("Header to send (%s): %x\n", command, header)

	if _, err := conn.Write(header); err != nil {
		return err
	}
	if _, err := conn.Write(payloadBytes); err != nil {
		return err
	}

	return nil
}

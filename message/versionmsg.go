package message

import (
	"bytes"
	"encoding/binary"
	"net"
	"time"
)

// VersionMessage represents a Bitcoin version message
type VersionMessage struct {
	Version     uint32
	Services    uint64
	Timestamp   int64
	AddrRecv    NetAddr
	AddrFrom    NetAddr
	Nonce       uint64
	UserAgent   string
	StartHeight int32
	Relay       bool
}

// NetAddr represents a network address in the Bitcoin protocol
type NetAddr struct {
	Services uint64
	IP       [16]byte
	Port     uint16
}

// createVersionMessage creates a version message to send to the target node
func CreateVersionMessage(targetNode string) VersionMessage {
	var addrRecv NetAddr
	copy(addrRecv.IP[:], net.ParseIP(targetNode).To16())
	addrRecv.Port = binary.BigEndian.Uint16([]byte{0x20, 0x8d}) // 8333 in Big Endian

	return VersionMessage{
		Version:     ProtocolVersion,
		Services:    Services,
		Timestamp:   time.Now().Unix(),
		AddrRecv:    addrRecv,
		AddrFrom:    NetAddr{},
		Nonce:       uint64(time.Now().UnixNano()),
		UserAgent:   UserAgent,
		StartHeight: 0,
		Relay:       false,
	}
}

// encodeVersionMessage encodes a VersionMessage into bytes
func EncodeVersionMessage(msg VersionMessage) ([]byte, error) {
	var buf bytes.Buffer

	// Write fixed-size fields
	if err := binary.Write(&buf, binary.LittleEndian, msg.Version); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, msg.Services); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, msg.Timestamp); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, msg.AddrRecv); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, msg.AddrFrom); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, msg.Nonce); err != nil {
		return nil, err
	}

	// Write variable-size fields
	userAgentBytes := []byte(msg.UserAgent)
	if err := buf.WriteByte(byte(len(userAgentBytes))); err != nil {
		return nil, err
	}
	if _, err := buf.Write(userAgentBytes); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, msg.StartHeight); err != nil {
		return nil, err
	}
	if err := buf.WriteByte(boolToByte(msg.Relay)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// decodeVersionMessage decodes bytes into a VersionMessage
func DecodeVersionMessage(payload []byte) (VersionMessage, error) {
	var msg VersionMessage
	buf := bytes.NewReader(payload)

	// Read fixed-size fields
	if err := binary.Read(buf, binary.LittleEndian, &msg.Version); err != nil {
		return msg, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.Services); err != nil {
		return msg, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.Timestamp); err != nil {
		return msg, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.AddrRecv); err != nil {
		return msg, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.AddrFrom); err != nil {
		return msg, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.Nonce); err != nil {
		return msg, err
	}

	// Read variable-size fields
	var userAgentLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &userAgentLen); err != nil {
		return msg, err
	}
	userAgentBytes := make([]byte, userAgentLen)
	if _, err := buf.Read(userAgentBytes); err != nil {
		return msg, err
	}
	msg.UserAgent = string(userAgentBytes)

	if err := binary.Read(buf, binary.LittleEndian, &msg.StartHeight); err != nil {
		return msg, err
	}
	var relayByte uint8
	if err := binary.Read(buf, binary.LittleEndian, &relayByte); err != nil {
		return msg, err
	}
	msg.Relay = relayByte == 1

	return msg, nil
}

// boolToByte converts a bool to a byte
func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

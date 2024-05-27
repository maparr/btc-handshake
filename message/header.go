package message

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/maparr/btc-handshake/checksum"
)

const (
	StartString = "f9beb4d9" // Mainnet magic value
)

func CreateMessageHeader(command string, payload []byte) []byte {
	var buf bytes.Buffer

	magic, _ := hex.DecodeString(StartString)
	buf.Write(magic)

	cmd := make([]byte, 12)
	copy(cmd, command)
	buf.Write(cmd)

	length := uint32(len(payload))
	binary.Write(&buf, binary.LittleEndian, length)

	checksum := checksum.CreateChecksum(payload)
	buf.Write(checksum[:])

	return buf.Bytes()
}
